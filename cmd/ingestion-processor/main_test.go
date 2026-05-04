package main

import "testing"

func TestProcessRecords_ValidInput(t *testing.T) {

	lines := []string{
		"john,25,engineer",
		"sarah,30,manager",
		
	}
	
	records, errors := processRecords(lines)

	if len(records) !=2 {
		t.Errorf("expected 2 records, got %d", len(records))
	}

	if len(errors) !=0 {
		t.Errorf("expected 0 errors, got %d", len(errors))
	}
	
}

func TestProcessRecords_InvalidData(t *testing.T) {

	lines := []string{
		"john,25,engineer",
		"bad,abc,test",
		"invalid,line",
	}
	records, errors := processRecords(lines)

	if len(records) != 1{
		t.Errorf("expected 1 valid record, got %d", len(records))
	}

	if len(errors) != 2{
		t.Errorf("expected 2 errors, got %d", len(errors))
	}
}

