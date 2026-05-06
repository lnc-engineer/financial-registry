concept

systems should return a single structured response combining data and errors

goal is to unify outputs into one JSON object for API-style communication

How it works

input: []Record, []string
process: wrap into Response struct
output: single JSON object

TEST

valid records -> included in response.records  
errors -> included in response.errors  
output -> structured JSON  

code pattern

type Response struct {
    Records []Record `json:"records"`
    Errors  []string `json:"errors"`
}

response := Response{
    Records: validRecords,
    Errors:  errors,
}

jsonData, _ := json.MarshalIndent(response, "", " ")

System connection

processing engine → response struct → JSON output

system now behaves like an API response layer

foundation for web services and microservices
