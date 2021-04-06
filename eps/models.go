package main

// Company ...
type Company struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// S3Addr ...
type S3Addr struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type Event struct {
	Address S3Addr  `json:"address"`
	Company Company `json:"company"`
}

// EPSData ...
type EPSData struct {
	Year int     `json:"year"`
	EPS  float64 `json:"eps"`
}
