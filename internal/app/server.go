package app

import "net/http"

// NewHTTPServer creates an HTTP server for the registry application.
func NewHTTPServer(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: NewHTTPHandler(),
	}
}

// NewConfiguredHTTPServer creates an HTTP server using application configuration.
func NewConfiguredHTTPServer(cfg Config) *http.Server {
	return NewHTTPServer(cfg.HTTPAddr)
}
