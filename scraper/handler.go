package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Handle ...
func Handle(request Request) (response Response, err error) {
	for page := 1; page <= nPages; page++ {
		url := toURL(request.Company, page)
		pageBody, err := scrapePage(url)
		if err != nil {
			return response, err
		}
		defer pageBody.Close()

		addr := toS3Addr(request.Company, page)
		if err := uploadPage(addr, pageBody); err != nil {
			return response, err
		}

		// event := toEvent(addr, request.Company)
		// if err := sendEvent(event); err != nil {
		// 	return response, err
		// }
		// EventBridge to Kinesis isn't yet supported by localstack
		// See: https://github.com/localstack/localstack/issues/3826
		// Put record directly to Kinesis for now
		event := toEvent(addr, request.Company)
		if err := sendEventToKenesis(event); err != nil {
			return response, err
		}
	}
	return Response{Success: true}, nil
}

func scrapePage(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status: %s", resp.Status)
	}
	return resp.Body, nil
}

func uploadPage(addr S3Addr, page io.Reader) error {
	sess, err := session.NewSession(&awsConfig)
	if err != nil {
		return err
	}

	uploader := s3manager.NewUploader(sess)

	input := &s3manager.UploadInput{
		Bucket: &addr.Bucket,
		Key:    &addr.Key,
		Body:   page,
	}
	_, err = uploader.Upload(input)
	return err
}

func sendEvent(event Event) error {
	bytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	sess, err := session.NewSession(&awsConfig)
	if err != nil {
		return err
	}

	svc := eventbridge.New(sess)

	_, err = svc.PutEvents(&eventbridge.PutEventsInput{
		Entries: []*eventbridge.PutEventsRequestEntry{
			{
				Detail:       aws.String(string(bytes)),
				EventBusName: aws.String(eventBusName),
				Source:       aws.String(scraperSource),
			},
		},
	})
	return err
}

func sendEventToKenesis(event Event) error {
	bytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	sess, err := session.NewSession(&awsConfig)
	if err != nil {
		return err
	}

	kc := kinesis.New(sess)

	_, err = kc.PutRecord(&kinesis.PutRecordInput{
		Data:         bytes,
		StreamName:   aws.String(streamName),
		PartitionKey: aws.String(event.Company.Symbol),
	})
	return err
}
