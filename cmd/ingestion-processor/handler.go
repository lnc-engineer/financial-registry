package main

import (
	"encoding/json"
	"net/http"

	"github.com/lnc-engineer/financial-registry/internal/execution"
)

// -------------------- RESPONSE WRITER --------------------

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// -------------------- JSON HELPER --------------------

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

// -------------------- HANDLER --------------------

func processHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"errors":  []string{"method not allowed"},
		})
		return
	}

	execCtx, ok := execution.FromContext(r.Context())
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

	response := ProcessIngestion(execCtx, request.Records)

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