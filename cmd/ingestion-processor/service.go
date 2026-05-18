package main

func ProcessIngestion(lines []string) Response {
	
	validRecords, errors := processRecords(lines)

	response := Response{
		Success: len(errors) == 0,
		Records: validRecords,
		Errors: errors,
	}

	return response
	}
