package main

import "fmt"

func toURL(company Company, index int) string {
	return fmt.Sprintf(urlTemplate, company.Name, company.Symbol, index)
}

func toS3Addr(company Company, index int) S3Addr {
	return S3Addr{
		Bucket: webpagesBucket,
		Key:    fmt.Sprintf(s3KeyTemplate, company.Name, company.Symbol, index),
	}
}

func toEvent(addr S3Addr, company Company) Event {
	return Event{
		Address: addr,
		Company: company,
	}
}
