package main

import (
	"fmt"
	"github.com/lnc-engineer/financial-registry/internal/execution"
	"strings"
	"time"
	"strconv"
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

func processRecords(ctx execution.ExecutionContext, lines []string) ([]Record, []string) {

	start := time.Now()

	rawRecords := parseLines(lines)

	var validRecords []Record
	var errorMessages []string

	for index, raw := range rawRecords {

    recordCtx := execution.NewChildSpan(ctx, "record-"+strconv.Itoa(index+1))
    execution.LogEvent(recordCtx, "record_processing_started")

    record, err := toRecord(raw)

    if err != nil {
        errorMessages = append(errorMessages,
            fmt.Sprintf("Invalid record at line %d: %s", index+1, raw.Raw))

        execution.RecordFailure()
        execution.LogEvent(recordCtx, "record_failed")
        continue
    }

    validRecords = append(validRecords, record)

    execution.RecordSuccess()
    execution.LogEvent(recordCtx, "record_success")
}

	fmt.Println("[METRICS] processing duration:", time.Since(start))

	return validRecords, errorMessages
}
