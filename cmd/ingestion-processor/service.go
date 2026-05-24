package main

func ProcessIngestion(lines []string) ProcessResponse {

	validRecords, errors := processRecords(lines)

	response := ProcessResponse{
		Success: len(errors) == 0,
		Records: validRecords,
		Errors:  errors,
	}

	return response
}
