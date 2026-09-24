package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewHTTPHandler_CreateAndGetSystem(t *testing.T) {
	handler := NewHTTPHandler()

	createBody := `{
		"id": "system-1",
		"name": "Payments System",
		"description": "Core payments platform",
		"status": "active"
	}`

	createRequest := httptest.NewRequest(
		http.MethodPost,
		"/systems",
		strings.NewReader(createBody),
	)
	createRecorder := httptest.NewRecorder()

	handler.ServeHTTP(createRecorder, createRequest)

	if createRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, createRecorder.Code)
	}

	getRequest := httptest.NewRequest(
		http.MethodGet,
		"/systems/system-1",
		nil,
	)
	getRecorder := httptest.NewRecorder()

	handler.ServeHTTP(getRecorder, getRequest)

	if getRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, getRecorder.Code)
	}

	if !strings.Contains(getRecorder.Body.String(), `"Payments System"`) {
		t.Fatalf("expected response to contain Payments System, got %q", getRecorder.Body.String())
	}
}
