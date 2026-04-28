concept

raw lines are not useful until they become structured records

goal is to transform string lines into typed/structured data the system can reason about

How it works

input: []string (lines from parser)
process: map each line → structured Record object
output: []Record (clean system-ready format)

TEST valid lines -> parsed into records malformed lines -> skipped or flagged empty lines -> ignored

code pattern

type Record struct {
    Raw string
    Fields []string
}

records := make([]Record, 0)

for _, line := range lines {
    if strings.TrimSpace(line) == "" {
        continue
    }

    record := Record{
        Raw: line,
        Fields: strings.Fields(line),
    }

    records = append(records, record)
}

System connection

parser → structuring layer

system now moves from raw text ingestion to structured data representation

foundation for validation, filtering, and future indexing in registry pipeline
