package main

import (
	"github.com/lnc-engineer/financial-registry/internal/execution"
	"strconv"
)

func ProcessIngestion(ctx execution.ExecutionContext, records []RawRecord) ProcessResponse {

	execution.LogEvent(ctx, "ingestion_started")

	lines := make([]string, 0, len(records))
	for _, r := range records {
		lines = append(lines, r.Raw)
	}

	execution.LogEvent(ctx, "records_received")

	validRecords, errors := processRecords(lines)

	ctx.Metadata["records_processed"] =
		strconv.Itoa(len(validRecords))

	execution.LogEvent(ctx, "records_processed")

	resp := ProcessResponse{
		Success: len(errors) == 0,
		Records: validRecords,
		Errors:  errors,
	}

	execution.LogEvent(ctx, "ingestion_completed")

	return resp
}
