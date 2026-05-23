package main

import "testing"

func TestProcessRecords_MixedInput(t *testing.T) {

	lines := []string{
		"john,25,engineer",
		"bad,abc,test",
		"sarah,30,manager",
	}

	records, errors := processRecords(lines)

	if len(records) != 2 {
		t.Errorf("expected 2 valid records, got %d", len(records))
	}

	if len(errors) != 1 {
		t.Errorf("expected 1 error, got %d", len(errors))
	}
	if len(errors) > 0 && errors[0] != "Invalid age at line 2: bad,abc,test" {
		t.Errorf("unexpected error message: %s", errors[0])
	}
}
func TestProcessRecords_InvalidData(t *testing.T) {

	lines := []string{
		"john,25,engineer",
		"bad,abc,test",
		"invalid,line",
	}
	records, errors := processRecords(lines)

	if len(records) != 1 {
		t.Errorf("expected 1 valid record, got %d", len(records))
	}

	if len(errors) != 2 {
		t.Errorf("expected 2 errors, got %d", len(errors))
	}
}
