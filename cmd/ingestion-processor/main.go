package main

import (
	"fmt"
	"strconv"
	"strings"
	"encoding/json"
	"net/http"
	"io"
)

// Generic structure
type RawRecord struct {
	Raw    string   // original line
	Fields []string // split parts
}

// Final structured record
type Record struct {
	Name string
	Age  int
	Role string
}

// Convert lines -> RawRecord
func parseLines(lines []string) []RawRecord {
	records := make([]RawRecord, 0)

	for _, line := range lines {

		// ignore empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		record := RawRecord{
			Raw:    strings.TrimSpace(line),
			Fields: strings.Split(strings.TrimSpace(line), ","),
		}

		records = append(records, record)
	}

	return records
}

// Convert RawRecord -> Record
func toRecord(r RawRecord) (Record, error) {

	if len(r.Fields) != 3 {
		return Record{}, fmt.Errorf("invalid field count: expected 3")
	}

	// convert age to int
	age, err := strconv.Atoi(strings.TrimSpace(r.Fields[1]))
	if err != nil {
		return Record{}, fmt.Errorf("invalid age")
	}

	return Record{
		Name: strings.TrimSpace(r.Fields[0]),
		Age:  age,
		Role: strings.TrimSpace(r.Fields[2]),
	}, nil
}

func processRecords(lines []string) ([]Record, []string) {

	rawRecords := parseLines(lines)

	var validRecords []Record
	var errorMessages []string

	for index, raw := range rawRecords {
		record, err := toRecord(raw)

		if err != nil {

			if err.Error() == "invalid age" {
				errorMessages = append(errorMessages,
					fmt.Sprintf("Invalid age at line %d: %s", index+1, raw.Raw))
			} else {
				errorMessages = append(errorMessages,
					fmt.Sprintf("Invalid record at line %d: %s", index+1, raw.Raw))
			}

			continue
		}

		validRecords = append(validRecords, record)
	}

	return validRecords, errorMessages
}

type Response struct {
	Records []Record `json:"records"`
	Errors []string `json:"errors"`
}


func processHandler(w http.ResponseWriter, r *http.Request)  {

	if r.Method != http.MethodPost {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	lines := strings.Split(string(body), "\n")

	validRecords, errors := processRecords(lines)

	response := Response{
		Records: validRecords,
		Errors: errors,
	}

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func main() {

	http.HandleFunc("/process", processHandler)

	fmt.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}


/*func main() {
	fmt.Println("Ingestion Processor Started")

	if len(os.Args) < 2 {
		fmt.Println("No input file provided")
		return
	}

	for i := 1; i < len(os.Args); i++ {
		file := os.Args[i]

		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		fmt.Println("Processing file:", file)

		lines := strings.Split(string(data), "\n")

		validRecords, errors := processRecords(lines)

		response := Response{
			Records: validRecords,
			Errors: errors,
		}

		fmt.Println("\n--- Processing Summary ---")
		fmt.Printf("Total valid records: %d\n", len(validRecords))
		fmt.Printf("Total errors: %d\n\n", len(errors))


		jsonData, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Println("Error converting to JSON:", err)
			return
		}

		fmt.Println("\nFinal Output (JSON):")
		fmt.Println(string(jsonData))
		
	}
}*/
