package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
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

		key := toS3Key(request.Company, page)
		if err := uploadPage(key, pageBody); err != nil {
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

func uploadPage(key string, page io.Reader) error {
	sess, err := session.NewSession(&awsConfig)
	if err != nil {
		return err
	}

	uploader := s3manager.NewUploader(sess)

	input := &s3manager.UploadInput{
		Bucket: &webpagesBucket,
		Key:    &key,
		Body:   page,
	}
	_, err = uploader.Upload(input)
	return err
}
