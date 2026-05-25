package main

func ProcessIngestion(records []RawRecord) ProcessResponse {

	lines := make([]string, 0, len(records))
	for _, r := range records {
		lines = append(lines, r.Raw)
	}

	validRecords, errors := processRecords(lines)

	return ProcessResponse{
		Success: len(errors) == 0,
		Records: validRecords,
		Errors:  errors,
	}
}
