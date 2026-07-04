package main

import (
	"github.com/lnc-engineer/financial-registry/internal/execution"
	"strconv"
)

func ProcessIngestion(ctx execution.ExecutionContext, records []string) ProcessResponse {

	execution.LogEvent(ctx, "ingestion_started")

	lines := records

	execution.LogEvent(ctx, "records_received")

	validRecords, errors := processRecords(ctx, lines)

	ctx = ctx.WithAttribute(
		"records_processed",
		strconv.Itoa(len(validRecords)),
	)

	execution.LogEvent(ctx, "records_processed")

	resp := ProcessResponse{
		Success: len(errors) == 0,
		Records: validRecords,
		Errors:  errors,
	}

	execution.LogEvent(ctx, "ingestion_completed")

	return resp
}
