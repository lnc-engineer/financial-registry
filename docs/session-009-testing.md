concept

systems must be tested to ensure correctness and reliability

goal is to verify that processing logic behaves as expected for different inputs

How it works

input: []string (test data)
process: run processRecords function
output: compare actual results with expected results

TEST

valid input -> correct number of records, no errors  
mixed input -> valid records processed, invalid ones flagged  
invalid input -> no valid records, errors returned  

code pattern

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

System connection

processing layer → tested with controlled inputs

system now ensures correctness before handling real data

foundation for reliability, debugging, and safe system evolution
