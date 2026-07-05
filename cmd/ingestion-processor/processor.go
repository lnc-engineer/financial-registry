package main

import (
	"fmt"
	"github.com/lnc-engineer/financial-registry/internal/execution"
	"strconv"
	"strings"
	"time"
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

	parts := strings.Split(strings.TrimSpace(r.Raw), ",")

	if len(parts) != 3 {
		return Record{}, fmt.Errorf("invalid format")
	}

	name := strings.TrimSpace(parts[0])
	ageStr := strings.TrimSpace(parts[1])
	role := strings.TrimSpace(parts[2])

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return Record{}, fmt.Errorf("invalid age")
	}

	return Record{
		Name: name,
		Age:  age,
		Role: role,
	}, nil
}

func processRecords(ctx execution.ExecutionContext, lines []string) ([]Record, []string) {

	start := time.Now()

	rawRecords := parseLines(lines)

	var validRecords []Record
	var errorMessages []string

	for index, raw := range rawRecords {

		recordCtx := execution.NewChildSpan(ctx, "record-"+strconv.Itoa(index+1))
		recordCtx = execution.StartSpan(recordCtx)

		recordCtx = recordCtx.WithAttribute("stage", "record_processing")
		recordCtx = recordCtx.WithAttribute("record_index", strconv.Itoa(index+1))

		execution.LogEvent(recordCtx, "record_processing_started")

		record, err := toRecord(raw)

		if err != nil {
			errorMessages = append(errorMessages,
				fmt.Sprintf("Invalid record at line %d: %s", index+1, raw.Raw))

			recordCtx = recordCtx.WithAttribute("result", "failure")
			recordCtx.Status = "failure"

			execution.RecordFailure(recordCtx)
			execution.LogEvent(recordCtx, "record_failed")

			recordCtx = execution.FinishSpan(recordCtx)
			continue
		}

		validRecords = append(validRecords, record)

		recordCtx = recordCtx.WithAttribute("result", "success")
		recordCtx.Status = "success"

		execution.RecordSuccess(recordCtx)
		execution.LogEvent(recordCtx, "record_success")

		recordCtx = execution.FinishSpan(recordCtx)
	}

	fmt.Println("[METRICS] processing duration:", time.Since(start))

	return validRecords, errorMessages

}
