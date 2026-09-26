package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lnc-engineer/financial-registry/internal/app"
)

const shutdownTimeout = 5 * time.Second

func main() {
	server := app.NewHTTPServer(":8080")

	if err := runServer(server); err != nil {
		log.Fatal(err)
	}
}

func runServer(server *http.Server) error {
	log.Printf("registry server listening on %s", server.Addr)

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(signalChannel)

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil

	case signalValue := <-signalChannel:
		log.Printf("received signal %s, shutting down", signalValue)

		shutdownContext, cancel := context.WithTimeout(
			context.Background(),
			shutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownContext); err != nil {
			return err
		}

		log.Println("registry server stopped")
		return nil
	}
}
