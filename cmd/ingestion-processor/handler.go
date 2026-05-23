package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

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

	lines := request.Lines
	response := ProcessIngestion(lines)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	status := http.StatusOK
	if !response.Success {
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

		start := time.Now() // start timer

		requestCount++

		reqID := generateRequestID()

		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next(rw, r)

		duration := time.Since(start) // calculate duration

		fmt.Printf(
			"[%s] [%s] %s | %d | %v | Total Requests: %d\n",
			reqID,
			r.Method,
			r.URL.Path,
			rw.statusCode,
			duration,
			requestCount,
		)

	}
}
