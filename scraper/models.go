package main

// Company ...
type Company struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// Request ...
type Request struct {
	Company Company `json:"company"`
}

// Respnose ...
type Response struct {
	Success bool `json:"success"`
}

// S3Addr ...
type S3Addr struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type Event struct {
	Address S3Addr
	Company Company
}
