package main

import (
	"net/http"
	"testing"
	"time"
)

func TestRunServerReturnsWhenServerIsClosed(t *testing.T) {
	server := &http.Server{
		Addr: ":0",
	}

	result := make(chan error, 1)

	go func() {
		result <- runServer(server)
	}()

	time.Sleep(50 * time.Millisecond)

	if err := server.Close(); err != nil {
		t.Fatalf("close server: %v", err)
	}

	select {
	case err := <-result:
		if err != nil {
			t.Fatalf("runServer returned error: %v", err)
		}

	case <-time.After(time.Second):
		t.Fatal("runServer did not return after server close")
	}
}
