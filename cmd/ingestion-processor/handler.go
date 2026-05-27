package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/lnc-engineer/financial-registry/internal/execution"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}



type contextKey string



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

func ProcessRecords(execCtx execution.ExecutionContext, request ProcessRequest) ProcessResponse {

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

	execCtx, ok := r.Context().Value(execution.ExecutionContextKey).(execution.ExecutionContext)
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


