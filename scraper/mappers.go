package main

import "fmt"

func toURL(company Company, index int) string {
	return fmt.Sprintf(urlTemplate, company.Name, company.Symbol, index)
}

func toS3Key(company Company, index int) string {
	return fmt.Sprintf(s3KeyTemplate, company.Name, company.Symbol, index)
}
