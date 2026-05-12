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
