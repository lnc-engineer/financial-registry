concept

processing logic should not live inside main

goal is to extract core logic into a reusable function (engine layer) that can be used independently from input/output

How it works

input: []string (lines from file)
process:
- parse lines into RawRecord
- convert RawRecord → Record with validation
- collect valid records
- collect error messages
output:
- []Record (valid structured data)
- []string (error messages)

TEST

valid records -> returned in []Record 
invalid field count -> added to errors
invalid data type (e.g. age) -> added to errors
main only handles printing and file input

code pattern

func processRecords(lines []string) ([]Record, []string) {

    rawRecords := parseLines(lines)

    var validRecords []Record
    var errorMessages []string

    for index, raw := range rawRecords {
        record, err := toRecord(raw)

        if err != nil {

            if err.Error() == "invalid age" {
                errorMessages = append(errorMessages,
                    fmt.Sprintf("Invalid age at line %d: %s", index+1, raw.Raw))
            } else {
                errorMessages = append(errorMessages,
                    fmt.Sprintf("Invalid record at line %d: %s", index+1, raw.Raw))
            }

            continue
        }

        validRecords = append(validRecords, record)
    }

    return validRecords, errorMessages
}

main responsibilities

- read file
- split into lines
- call processRecords()
- display results

System connection

input layer (CLI/file)
    ↓
engine layer (processRecords)
    ↓
output layer (printing)

system now separates orchestration from processing logic

foundation for APIs, testing, and scalable system design
