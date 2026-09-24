package app

import "net/http"

// NewHTTPServer creates an HTTP server for the registry application.
func NewHTTPServer(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: NewHTTPHandler(),
	}
}
