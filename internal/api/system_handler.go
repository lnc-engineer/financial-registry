package api

import (
	"encoding/json"
	"net/http"

	"github.com/lnc-engineer/financial-registry/internal/registry"
)

// SystemHandler handles HTTP requests for registered systems.
type SystemHandler struct {
	service systemService
}

// NewSystemHandler creates a system HTTP handler using the provided service.
func NewSystemHandler(service systemService) *SystemHandler {
	return &SystemHandler{
		service: service,
	}
}

// Routes returns an HTTP handler with the system routes registered.
func (h *SystemHandler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /systems", h.CreateSystem)
	mux.HandleFunc("GET /systems", h.ListSystems)
	mux.HandleFunc("GET /systems/{id}", h.GetSystem)

	return mux
}

// createSystemRequest represents the JSON payload for creating a system.
type createSystemRequest struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Status      registry.SystemStatus `json:"status"`
}

// CreateSystem handles POST /systems requests.
func (h *SystemHandler) CreateSystem(w http.ResponseWriter, r *http.Request) {
	var request createSystemRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	system := registry.System{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		Status:      request.Status,
	}

	if err := h.service.Register(system); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetSystem handles GET /systems/{id} requests.
func (h *SystemHandler) GetSystem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	system, ok := h.service.Get(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(system); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

// ListSystems handles GET /systems requests.
func (h *SystemHandler) ListSystems(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	var systems []registry.System

	switch status {
	case "":
		systems = h.service.List()
	case string(registry.SystemStatusActive), string(registry.SystemStatusInactive):
		systems = h.service.ListByStatus(registry.SystemStatus(status))
	default:
		http.Error(w, "invalid status", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(systems); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
