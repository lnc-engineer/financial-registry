package main

import (
	"testing"

	"github.com/lnc-engineer/financial-registry/internal/execution"
)

func TestProcessRecords_MixedInput(t *testing.T) {

	lines := []string{
		"john,25,engineer",
		"bad,abc,test",
		"sarah,30,manager",
	}

	ctx := execution.ExecutionContext{}

	records, errors := processRecords(ctx, lines)

	if len(records) != 2 {
		t.Errorf("expected 2 valid records, got %d", len(records))
	}

	if len(errors) != 1 {
		t.Errorf("expected 1 error, got %d", len(errors))
	}
}

func TestProcessRecords_InvalidData(t *testing.T) {

	lines := []string{
		"john,25,engineer",
		"bad,abc,test",
		"invalid,line",
	}

	ctx := execution.ExecutionContext{}

	records, errors := processRecords(ctx, lines)

	if len(records) != 1 {
		t.Errorf("expected 1 valid record, got %d", len(records))
	}

	if len(errors) != 2 {
		t.Errorf("expected 2 errors, got %d", len(errors))
	}
}
