# Financial Registry

A Go-based ingestion and validation API designed to process
structured registry data through a layered backend pipeline.


## 🧠 Data Processing Pipeline
File (data.txt)
↓
Read File (os.ReadFile)
↓
Split into lines ([]string)
↓
parseLines()
↓
[]RawRecord
↓
toRecord()
↓
[]Record + errors
↓
processRecords()
↓
(validRecords, errorMessages)
↓
JSON Output (encoding/json)


### Explanation

- **Input Layer**: Reads raw file data
- **Parsing Layer**: Converts text → structured fields
- **Validation Layer**: Ensures data correctness
- **Processing Layer**: Separates valid records and errors
- **Output Layer**: Produces JSON (API-ready)

## Layers

- Input Layer
- Parsing Layer
- Structuring Layer
- Validation Layer
- Processing Engine
- Output Layer (JSON)

## Notes

This pipeline follows a simplified ETL pattern:
- Extract (file input)
- Transform (parsing + validation)
- Load (JSON output)

## Architecture

The ingestion processor follows a layered backend pipeline:

File Input / HTTP Request
↓
Request Parsing
↓
Validation
↓
Record Transformation
↓
Error Aggregation
↓
Structured JSON Response

Current project structure:

cmd/ingestion-processor/
├── main.go
├── handler.go
├── processor.go
└── models.go

Key concepts implemented:

- JSON APIs
- HTTP status codes
- request validation
- structured responses
- separation of concerns
- multi-file architecture
- automated testing

## Current Features

- Parse comma-separated records
- Validate structured input
- Detect invalid ages and malformed records
- Aggregate processing errors
- Return structured JSON responses
- Expose processing through HTTP API
- Handle HTTP status codes correctly
- Pretty-print JSON responses
- Multi-file backend architecture
- Unit testing with Go testing package

---

## Run The API

Start server:

```bash
go run ./cmd/ingestion-processor
```

Server runs on:

```text
http://localhost:8080
```

---

## Example Request

```bash
curl -i -X POST http://localhost:8080/process \
-H "Content-Type: application/json" \
-d '{
  "lines": [
    "john,25,engineer"
  ]
}'
```

---

## Example Success Response

```json
{
  "success": true,
  "records": [
    {
      "Name": "john",
      "Age": 25,
      "Role": "engineer"
    }
  ],
  "errors": null
}
```

---

## Example Error Response

```json
{
  "success": false,
  "records": null,
  "errors": [
    "Invalid age at line 1: bad,abc,test"
  ]
}
```
