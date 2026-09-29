package app

import (
	"testing"
	"time"
)

func TestNewHTTPServer(t *testing.T) {
	server := NewHTTPServer(":8080")

	if server.Addr != ":8080" {
		t.Fatalf("expected address %q, got %q", ":8080", server.Addr)
	}

	if server.Handler == nil {
		t.Fatal("expected server handler to be configured")
	}

	if server.ReadHeaderTimeout != 5*time.Second {
		t.Fatalf("expected read header timeout %s, got %s", 5*time.Second, server.ReadHeaderTimeout)
	}

	if server.ReadTimeout != 10*time.Second {
		t.Fatalf("expected read timeout %s, got %s", 10*time.Second, server.ReadTimeout)
	}

	if server.WriteTimeout != 10*time.Second {
		t.Fatalf("expected write timeout %s, got %s", 10*time.Second, server.WriteTimeout)
	}

	if server.IdleTimeout != 60*time.Second {
		t.Fatalf("expected idle timeout %s, got %s", 60*time.Second, server.IdleTimeout)
	}
}
