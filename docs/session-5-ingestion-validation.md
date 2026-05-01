concept

raw lines are not useful until they become structured records

goal is to transform string lines into structured and typed data the system can reason about, in two stages:
1. generic structuring (RawRecord)
2. domain mapping with validation (Record)

How it works

input: []string (lines from file)
process:
- map each line → RawRecord (split into fields)
- map RawRecord → Record (assign meaning + validate types)
- classify errors (invalid structure vs invalid data)
output: []Record (clean system-ready format)

TEST

valid lines -> parsed into records
invalid field count -> skipped 
invalid data type (e.g. age) -> flagged with specific error 
empty lines -> ignored 

code pattern

type RawRecord struct {
    Raw    string
    Fields []string
}

type Record struct {
    Name string
    Age  int
    Role string
}

rawRecords := make([]RawRecord, 0)

for _, line := range lines {
    if strings.TrimSpace(line) == "" {
        continue
    }

    raw := RawRecord{
        Raw:    line,
        Fields: strings.Split(strings.TrimSpace(line), ","),
    }

    rawRecords = append(rawRecords, raw)
}

for index, r := range rawRecords {

    if len(r.Fields) != 3 {
        fmt.Printf("Skipping invalid record (line %d): %s\n", index+1, r.Raw)
        continue
    }

    age, err := strconv.Atoi(strings.TrimSpace(r.Fields[1]))
    if err != nil {
        fmt.Printf("Invalid age at line %d: %s\n", index+1, r.Raw)
        continue
    }

    record := Record{
        Name: strings.TrimSpace(r.Fields[0]),
        Age:  age,
        Role: strings.TrimSpace(r.Fields[2]),
    }

    // use record
}

System connection

file reader → parser (lines) → structuring layer (RawRecord) → domain layer (Record + validation)

system now moves from raw text ingestion to structured, typed, and validated data

foundation for validation, filtering, transformation, error reporting, and future indexing in registry pipeline

