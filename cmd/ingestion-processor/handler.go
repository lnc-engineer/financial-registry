package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"math/rand"
	"strconv"
)

var requestCount int

func generateRequestID() string {
	return "REQ-" + strconv.Itoa(rand.Intn(100000))
}

func processHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request ProcessRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	validRecords, errors := processRecords(request.Lines)

	response := Response{
		Success: len(errors) == 0,
		Records: validRecords,
		Errors:  errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	status := http.StatusOK
	if len(errors) > 0 {
		status = http.StatusBadRequest
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(jsonData)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": "ok",
	}

	jsonData, _ := json.MarshalIndent(response, "", "  ")
	w.Write(jsonData)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"service": "financial-registry",
		"status":  "running",
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		requestCount++

		reqID := generateRequestID()

		fmt.Printf(
			"[%s] [%s] %s | Total Requests: %d\n",
			reqID,
			r.Method,
			r.URL.Path,
			requestCount,
		)

		next(w, r)
	}
}
