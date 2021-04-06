package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
)

const (
	localstackHostnameEnv = "LOCALSTACK_HOSTNAME"

	headersXPath   = "/html/body/section/div[2]/div/div[2]/div[2]/div/div[2]/div/div[1]/table/tbody/tr[1]/td"
	epsValuesXPath = "/html/body/section/div[2]/div/div[2]/div[2]/div/div[2]/div/div[1]/table/tbody/tr[5]/td"

	dateLayout = "Jan 06"
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
)
