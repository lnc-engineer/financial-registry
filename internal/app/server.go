package app

import (
	"net/http"
	"time"
)

const (
	httpReadHeaderTimeout = 5 * time.Second
	httpReadTimeout       = 10 * time.Second
	httpWriteTimeout      = 10 * time.Second
	httpIdleTimeout       = 60 * time.Second
)

// NewHTTPServer creates an HTTP server for the registry application.
func NewHTTPServer(addr string) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           NewHTTPHandler(),
		ReadHeaderTimeout: httpReadHeaderTimeout,
		ReadTimeout:       httpReadTimeout,
		WriteTimeout:      httpWriteTimeout,
		IdleTimeout:       httpIdleTimeout,
	}
}

// NewConfiguredHTTPServer creates an HTTP server using application configuration.
func NewConfiguredHTTPServer(cfg Config) *http.Server {
	return NewHTTPServer(cfg.HTTPAddr)
}
