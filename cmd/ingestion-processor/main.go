package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Generic structure
type RawRecord struct {
	Raw    string   // original line
	Fields []string // split parts
}

// Final structured record
type Record struct {
	Name string
	Age  int
	Role string
}

// Convert lines -> RawRecord
func parseLines(lines []string) []RawRecord {
	records := make([]RawRecord, 0)

	for _, line := range lines {

		// ignore empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		record := RawRecord{
			Raw:    line,
			Fields: strings.Split(strings.TrimSpace(line), ","),
		}

		records = append(records, record)
	}

	return records
}

// Convert RawRecord -> Record
func toRecord(r RawRecord) (Record, error) {

	if len(r.Fields) != 3 {
		return Record{}, fmt.Errorf("invalid field count: expected 3")
	}

	// convert age to int
	age, err := strconv.Atoi(strings.TrimSpace(r.Fields[1]))
	if err != nil {
		return Record{}, fmt.Errorf("invalid age")
	}

	return Record{
		Name: strings.TrimSpace(r.Fields[0]),
		Age:  age,
		Role: strings.TrimSpace(r.Fields[2]),
	}, nil
}

func main() {
	fmt.Println("Ingestion Processor Started")

	if len(os.Args) < 2 {
		fmt.Println("No input file provided")
		return
	}

	for i := 1; i < len(os.Args); i++ {
		file := os.Args[i]

		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		fmt.Println("Processing file:", file)

		lines := strings.Split(string(data), "\n")

		// Structuring layer
		rawRecords := parseLines(lines)
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
		fmt.Println("\n--- Processing Summary ---")
		fmt.Printf("Total valid records: %d\n", len(validRecords))
		fmt.Printf("Total errors: %d\n\n", len(errors))

		fmt.Println("Valid Records:")
		for _, r := range validRecords {
			fmt.Printf("%+v\n", r)
		}

		fmt.Println("\nErrors:")
		for _, e := range errors {
			fmt.Println(e)
		}
	}
}
