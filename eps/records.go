package main

// CompanyRecord ...
type CompanyRecord struct {
	Symbol string `dynamodbav:"Symbol"`
	Name   string `dynamodbav:"Name"`
}

// EPSRecord ...
type EPSRecord struct {
	Symbol string  `dynamodbav:"Symbol"`
	Year   int     `dynamodbav:"Year"`
	EPS    float64 `dynamodbav:"EPS"`
}
