package main

import (
	"fmt"
	"strings"
)

func parseLines(lines []string) []RawRecord {
	records := make([]RawRecord, 0)

	for _, line := range lines {

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		records = append(records, RawRecord{
			Raw:    line,
			Fields: strings.Split(line, ","), // optional future use
		})
	}

	return records
}

// UPDATED: now accepts simple records like "txn-001"
func toRecord(r RawRecord) (Record, error) {

	value := strings.TrimSpace(r.Raw)

	if value == "" {
		return Record{}, fmt.Errorf("empty record")
	}

	// Minimal valid transformation
	return Record{
		Name: value,
		Age:  0,
		Role: "unprocessed",
	}, nil
}

func processRecords(lines []string) ([]Record, []string) {

	rawRecords := parseLines(lines)

	var validRecords []Record
	var errorMessages []string

	for index, raw := range rawRecords {

		record, err := toRecord(raw)

		if err != nil {
			errorMessages = append(errorMessages,
				fmt.Sprintf("Invalid record at line %d: %s", index+1, raw.Raw))
			continue
		}

		validRecords = append(validRecords, record)
	}

	return validRecords, errorMessages
}