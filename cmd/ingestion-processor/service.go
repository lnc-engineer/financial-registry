package main

import (
	"github.com/lnc-engineer/financial-registry/internal/execution"
	"strconv"
)

func ProcessIngestion(ctx execution.ExecutionContext, records []string) ProcessResponse {

	execution.LogEvent(ctx, "ingestion_started")

	execution.LogEvent(ctx, "records_received")

	validRecords, errors := processRecords(ctx, records)

	// attach metadata to span
	ctx = ctx.WithAttribute(
		"records_processed",
		strconv.Itoa(len(validRecords)),
	)

	// set final status on span
	if len(errors) == 0 {
		ctx.Status = "success"
		execution.RecordSuccess(ctx)
	} else {
		ctx.Status = "failure"
		execution.RecordFailure(ctx)
	}

	execution.LogEvent(ctx, "records_processed")

	resp := ProcessResponse{
		Success: len(errors) == 0,
		Records: validRecords,
		Errors:  errors,
	}

	execution.LogEvent(ctx, "ingestion_completed")

	return resp
}
