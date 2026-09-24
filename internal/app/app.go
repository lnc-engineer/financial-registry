package app

import (
	"net/http"

	"github.com/lnc-engineer/financial-registry/internal/api"
	"github.com/lnc-engineer/financial-registry/internal/registry"
)

// NewHTTPHandler builds the registry application HTTP handler.
func NewHTTPHandler() http.Handler {
	store := registry.NewRegistry()
	service := registry.NewSystemService(store)
	handler := api.NewSystemHandler(service)

	return handler.Routes()
}
