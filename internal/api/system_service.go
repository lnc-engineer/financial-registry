package api

import "github.com/lnc-engineer/financial-registry/internal/registry"

// systemService defines the application operations required by the HTTP API.
type systemService interface {
	Register(system registry.System) error
	Get(id string) (registry.System, bool)
	List() []registry.System
	ListByStatus(status registry.SystemStatus) []registry.System
}

// Verify that the registry service satisfies the HTTP API service boundary.
var _ systemService = (*registry.SystemService)(nil)
