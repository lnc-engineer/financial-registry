package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseLines(lines []string) []RawRecord {
	records := make([]RawRecord, 0)

	for _, line := range lines {

		if strings.TrimSpace(line) == "" {
			continue
		}

		record := RawRecord{
			Raw:    strings.TrimSpace(line),
			Fields: strings.Split(strings.TrimSpace(line), ","),
		}

		records = append(records, record)
	}

	return records
}

func toRecord(r RawRecord) (Record, error) {

	if len(r.Fields) != 3 {
		return Record{}, fmt.Errorf("invalid field count: expected 3")
	}

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