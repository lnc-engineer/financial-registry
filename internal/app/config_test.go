package app

import "testing"

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.HTTPAddr != ":8080" {
		t.Fatalf("expected default HTTP address %q, got %q", ":8080", cfg.HTTPAddr)
	}
}

func TestNewConfiguredHTTPServer(t *testing.T) {
	cfg := Config{
		HTTPAddr: ":9090",
	}

	server := NewConfiguredHTTPServer(cfg)

	if server.Addr != ":9090" {
		t.Fatalf("expected server address %q, got %q", ":9090", server.Addr)
	}

	if server.Handler == nil {
		t.Fatal("expected server handler to be configured")
	}
}
