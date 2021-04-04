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
