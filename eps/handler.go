package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Handle ...
func Handle(ctx context.Context, kinesisEvent events.KinesisEvent) error {
	for _, record := range kinesisEvent.Records {
		if err := handleEventRecord(record); err != nil {
			return err
		}
	}
	return nil
}

func handleEventRecord(record events.KinesisEventRecord) error {
	var event Event
	if err := json.Unmarshal(record.Kinesis.Data, &event); err != nil {
		return err
	}

	pageBytes, err := downloadWebpage(event.Address)
	if err != nil {
		return err
	}

	epsData, err := getEPSData(pageBytes)
	if err != nil {
		return err
	}

	return fmt.Errorf("%+v", epsData)
}

func downloadWebpage(addr S3Addr) ([]byte, error) {
	sess, err := session.NewSession(&awsConfig)
	if err != nil {
		return nil, err
	}

	buffer := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(
		buffer,
		&s3.GetObjectInput{
			Bucket: aws.String(addr.Bucket),
			Key:    aws.String(addr.Key),
		},
	)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func getEPSData(pageBytes []byte) (epsData []EPSData, err error) {
	doc, err := htmlquery.Parse(bytes.NewReader(pageBytes))
	if err != nil {
		return
	}

	headers, err := htmlquery.QueryAll(doc, headersXPath)
	if err != nil {
		return
	}

	epsValues, err := htmlquery.QueryAll(doc, epsValuesXPath)
	if err != nil {
		return
	}

	// Skip first cell since it's row name
	for i := 1; i < len(headers) && i < len(epsValues); i++ {
		if htmlquery.SelectAttr(headers[i], "class") == "last_td" {
			continue
		}
		date, err := time.Parse(dateLayout, htmlquery.InnerText(headers[i]))
		if err != nil {
			return nil, err
		}
		eps, err := strconv.ParseFloat(htmlquery.InnerText(epsValues[i]), 64)
		if err != nil {
			return nil, err
		}
		epsData = append(
			epsData,
			EPSData{
				Year: date.Year(),
				EPS:  eps,
			},
		)
	}

	return
}
