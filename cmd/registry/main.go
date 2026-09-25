package main

import (
	"log"

	"github.com/lnc-engineer/financial-registry/internal/app"
)

func main() {
	server := app.NewHTTPServer(":8080")

	log.Printf("registry server listening on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
