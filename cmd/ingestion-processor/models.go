package main

type RawRecord struct {
	Raw    string
	Fields []string
}

type Record struct {
	Name string
	Age  int
	Role string
}

type ProcessResponse struct {
	Success bool     `json:"success"`
	Records []Record `json:"records"`
	Errors  []string `json:"errors"`
}

type ProcessRequest struct {
	Records []RawRecord `json:"records"`
}
