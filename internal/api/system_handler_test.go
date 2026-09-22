package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lnc-engineer/financial-registry/internal/registry"
)

func TestSystemHandler_CreateSystem(t *testing.T) {
	registryStore := registry.NewRegistry()
	service := registry.NewSystemService(registryStore)
	handler := NewSystemHandler(service)

	body := `{
		"id": "system-1",
		"name": "Payments System",
		"description": "Core payments platform",
		"status": "active"
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/systems",
		strings.NewReader(body),
	)

	recorder := httptest.NewRecorder()

	handler.CreateSystem(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, recorder.Code)
	}

	system, ok := service.Get("system-1")
	if !ok {
		t.Fatal("expected system to be registered")
	}

	if system.Name != "Payments System" {
		t.Fatalf("expected system name %q, got %q", "Payments System", system.Name)
	}
}

func TestSystemHandler_CreateSystem_InvalidJSON(t *testing.T) {
	registryStore := registry.NewRegistry()
	service := registry.NewSystemService(registryStore)
	handler := NewSystemHandler(service)

	request := httptest.NewRequest(
		http.MethodPost,
		"/systems",
		strings.NewReader(`{"id":`),
	)

	recorder := httptest.NewRecorder()

	handler.CreateSystem(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestSystemHandler_CreateSystem_ValidationError(t *testing.T) {
	registryStore := registry.NewRegistry()
	service := registry.NewSystemService(registryStore)
	handler := NewSystemHandler(service)

	body := `{
		"id": "",
		"name": "Payments System",
		"description": "Core payments platform",
		"status": "active"
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/systems",
		strings.NewReader(body),
	)

	recorder := httptest.NewRecorder()

	handler.CreateSystem(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestSystemHandler_GetSystem(t *testing.T) {
	registryStore := registry.NewRegistry()
	service := registry.NewSystemService(registryStore)
	handler := NewSystemHandler(service)

	err := service.Register(registry.System{
		ID:          "system-1",
		Name:        "Payments System",
		Description: "Core payments platform",
		Status:      registry.SystemStatusActive,
	})
	if err != nil {
		t.Fatalf("failed to register test system: %v", err)
	}

	request := httptest.NewRequest(
		http.MethodGet,
		"/systems/system-1",
		nil,
	)
	request.SetPathValue("id", "system-1")

	recorder := httptest.NewRecorder()

	handler.GetSystem(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("expected content type %q, got %q", "application/json", contentType)
	}

	expected := `"Payments System"`
	if !strings.Contains(recorder.Body.String(), expected) {
		t.Fatalf("expected response to contain %q, got %q", expected, recorder.Body.String())
	}
}

func TestSystemHandler_GetSystem_NotFound(t *testing.T) {
	registryStore := registry.NewRegistry()
	service := registry.NewSystemService(registryStore)
	handler := NewSystemHandler(service)

	request := httptest.NewRequest(
		http.MethodGet,
		"/systems/missing",
		nil,
	)
	request.SetPathValue("id", "missing")

	recorder := httptest.NewRecorder()

	handler.GetSystem(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, recorder.Code)
	}
}

func TestSystemHandler_Routes(t *testing.T) {
	registryStore := registry.NewRegistry()
	service := registry.NewSystemService(registryStore)
	handler := NewSystemHandler(service)

	request := httptest.NewRequest(
		http.MethodGet,
		"/systems/missing",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, recorder.Code)
	}
}

func TestSystemHandler_ListSystems(t *testing.T) {
	registryStore := registry.NewRegistry()
	service := registry.NewSystemService(registryStore)
	handler := NewSystemHandler(service)

	err := service.Register(registry.System{
		ID:          "system-1",
		Name:        "Payments System",
		Description: "Core payments platform",
		Status:      registry.SystemStatusActive,
	})
	if err != nil {
		t.Fatalf("failed to register first test system: %v", err)
	}

	err = service.Register(registry.System{
		ID:          "system-2",
		Name:        "Trading System",
		Description: "Trading platform",
		Status:      registry.SystemStatusInactive,
	})
	if err != nil {
		t.Fatalf("failed to register second test system: %v", err)
	}

	request := httptest.NewRequest(
		http.MethodGet,
		"/systems",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	body := recorder.Body.String()

	if !strings.Contains(body, `"Payments System"`) {
		t.Fatalf("expected response to contain Payments System, got %q", body)
	}

	if !strings.Contains(body, `"Trading System"`) {
		t.Fatalf("expected response to contain Trading System, got %q", body)
	}
}

func TestSystemHandler_ListSystemsByStatus(t *testing.T) {
	registryStore := registry.NewRegistry()
	service := registry.NewSystemService(registryStore)
	handler := NewSystemHandler(service)

	err := service.Register(registry.System{
		ID:          "system-1",
		Name:        "Payments System",
		Description: "Core payments platform",
		Status:      registry.SystemStatusActive,
	})
	if err != nil {
		t.Fatalf("failed to register first test system: %v", err)
	}

	err = service.Register(registry.System{
		ID:          "system-2",
		Name:        "Trading System",
		Description: "Trading platform",
		Status:      registry.SystemStatusInactive,
	})
	if err != nil {
		t.Fatalf("failed to register second test system: %v", err)
	}

	request := httptest.NewRequest(
		http.MethodGet,
		"/systems?status=active",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	body := recorder.Body.String()

	if !strings.Contains(body, `"Payments System"`) {
		t.Fatalf("expected response to contain Payments System, got %q", body)
	}

	if strings.Contains(body, `"Trading System"`) {
		t.Fatalf("expected response not to contain Trading System, got %q", body)
	}
}

func TestSystemHandler_ListSystemsByStatus_InvalidStatus(t *testing.T) {
	registryStore := registry.NewRegistry()
	service := registry.NewSystemService(registryStore)
	handler := NewSystemHandler(service)

	request := httptest.NewRequest(
		http.MethodGet,
		"/systems?status=unknown",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.Routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}
