concept

raw lines are not useful until they become structured records

goal is to transform string lines into structured data the system can reason about, in two stages:
1. generic structuring (RawRecord)
2. domain mapping (Record)

How it works

input: []string (lines from file)
process:
- map each line → RawRecord (split into fields)
- map RawRecord → Record (assign meaning + validate)
output: []Record (clean system-ready format)

TEST

valid lines -> parsed into records 
malformed lines -> skipped 
empty lines -> ignored 

code pattern

type RawRecord struct {
    Raw    string
    Fields []string
}

type Record struct {
    Name string
    Age  string
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

for _, r := range rawRecords {
    if len(r.Fields) != 3 {
        continue
    }

    record := Record{
        Name: strings.TrimSpace(r.Fields[0]),
        Age:  strings.TrimSpace(r.Fields[1]),
        Role: strings.TrimSpace(r.Fields[2]),
    }

    // use record
}

System connection

file reader → parser (lines) → structuring layer (RawRecord) → domain layer (Record)

system now moves from raw text ingestion to structured, validated data

foundation for validation, filtering, transformation, and future indexing in registry pipeline
