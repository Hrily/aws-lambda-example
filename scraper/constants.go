package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
)

const (
	urlTemplate = "https://www.moneycontrol.com/financials/%s/consolidated-ratiosVI/%s/%d"

	s3KeyTemplate = "%s/%s/%d"

	// nPages to scrape
	nPages = 3

	localstackHostnameEnv = "LOCALSTACK_HOSTNAME"
)

var (
	localstackS3Endpoint = fmt.Sprintf("http://%s:4566", os.Getenv(localstackHostnameEnv))

	awsConfig = aws.Config{
		Credentials:                   credentials.NewStaticCredentials("test", "test", ""),
		Endpoint:                      aws.String(localstackS3Endpoint),
		Region:                        aws.String(endpoints.UsEast1RegionID),
		S3ForcePathStyle:              aws.Bool(true),
		CredentialsChainVerboseErrors: aws.Bool(true),
	}

	webpagesBucket = "webpages"
)
