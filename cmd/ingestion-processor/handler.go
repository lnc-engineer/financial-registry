package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"sync/atomic"
	"context"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

type ExecutionContext struct {
	RequestID string
	StartTime time.Time
}

type contextKey string

const executionContextKey contextKey = "execution_context"


var requestCount int64

func generateRequestID() string {
	return "REQ-" + strconv.Itoa(rand.Intn(100000))
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}


func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	jsonData, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// service layer

func ProcessRecords(execCtx ExecutionContext, request ProcessRequest) ProcessResponse {

	fmt.Printf(
		"[%s] Processing %d records\n",
		execCtx.RequestID,
		len(request.Records),
	)

	return ProcessResponse{
		Success: true,
		Records: []Record{},
		Errors:  nil,
	}

}

//Handler

func processHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"errors":  []string{"method not allowed"},
		})
		return
	}

	execCtx, ok := r.Context().Value(executionContextKey).(ExecutionContext)
	if !ok {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"errors":  []string{"execution context missing"},
	})
	return
}

	defer r.Body.Close()

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
		writeJSON(w, http.StatusBadRequest, ProcessResponse{
			Success: false,
			Records: []Record{},
			Errors:  []string{"no records provided"},
		})
		return
	}

	response := ProcessRecords(execCtx, request)

	writeJSON(w, http.StatusOK, response)
}

//other handlers

func healthHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"errors":  []string{"method not allowed"},
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	writeJSON(w, http.StatusOK, map[string]string{
		"service": "financial-registry",
		"status":  "running",
	})

}


// middleware

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		atomic.AddInt64(&requestCount, 1)

		reqID := generateRequestID()

		execCtx := ExecutionContext{
		RequestID: reqID,
		StartTime: start,
	}

		ctx := context.WithValue(r.Context(), executionContextKey, execCtx)

		r = r.WithContext(ctx)

		w.Header().Set("X-Request-ID", reqID)

		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next(rw, r)

		if rw.statusCode == 0 {
			rw.statusCode = http.StatusOK
		}

		duration := time.Since(start)


fmt.Printf(
	`{"request_id":"%s","method":"%s","path":"%s","status":%d,"duration":"%v","total_requests":%d}`+"\n",
	reqID,
	r.Method,
	r.URL.Path,
	rw.statusCode,
	duration,
	atomic.LoadInt64(&requestCount),
)
		
	}
}