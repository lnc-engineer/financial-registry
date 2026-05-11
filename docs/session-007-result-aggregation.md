concept

processing should not immediately print results

goal is to separate data processing from output by collecting results first, then presenting them in a controlled way

How it works

input: []RawRecord (structured input)
process:
- attempt to convert each RawRecord → Record
- store valid records
- store errors separately with context
output:
- []Record (valid data)
- []string (error messages)
- summary of processing

TEST

valid records -> collected into validRecords
invalid field count -> added to errors
invalid data type (e.g. age) -> added to errors with specific message
empty lines -> ignored

code pattern

var validRecords []Record
var errors []string

for index, raw := range rawRecords {
    record, err := toRecord(raw)

    if err != nil {

        if err.Error() == "invalid age" {
            errors = append(errors,
                fmt.Sprintf("Invalid age at line %d: %s", index+1, raw.Raw))
        } else {
            errors = append(errors,
                fmt.Sprintf("Invalid record at line %d: %s", index+1, raw.Raw))
        }

        continue
    }

    validRecords = append(validRecords, record)
}

fmt.Println("--- Processing Summary ---")
fmt.Printf("Total valid records: %d\n", len(validRecords))
fmt.Printf("Total errors: %d\n", len(errors))

System connection

file reader → parser → structuring layer → validation layer → result aggregation → output layer

system now moves from immediate execution to controlled result handling

foundation for APIs, logging systems, and data pipelines
