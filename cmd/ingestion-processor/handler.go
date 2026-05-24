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


func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	jsonData, _ := json.MarshalIndent(payload, "", "  ")
	w.Write(jsonData)
}

// handlers

func processHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"errors":  []string{"method not allowed"},
		})
		return
	}

	var request ProcessRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"errors":  []string{"invalid JSON request"},
		})
		return
	}

	if len(request.Records) == 0 {
		response := ProcessResponse{
			Success: false,
			Records: nil,
			Errors:  []string{"no records provided"},
		}

		writeJSON(w, http.StatusBadRequest, response)
		return
	}

	// processing logic starts here

	response := ProcessResponse{
		Success: true,
		Records: nil,
		Errors:  nil,
	}

	writeJSON(w, http.StatusOK, response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"errors":  []string{"method not allowed"},
		})
		return
	}

	response := map[string]string{
		"status": "ok",
	}

	writeJSON(w, http.StatusOK, response)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"service": "financial-registry",
		"status":  "running",
	}

	writeJSON(w, http.StatusOK, response)
}


// middleware

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		requestCount++
		reqID := generateRequestID()

		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     0,
		}

		next(rw, r)

		if rw.statusCode == 0 {
			rw.statusCode = http.StatusOK
		}

		duration := time.Since(start)

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