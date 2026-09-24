package app

import "testing"

func TestNewHTTPServer(t *testing.T) {
	server := NewHTTPServer(":8080")

	if server.Addr != ":8080" {
		t.Fatalf("expected address %q, got %q", ":8080", server.Addr)
	}

	if server.Handler == nil {
		t.Fatal("expected server handler to be configured")
	}
}
