package main

import (
	"fmt"
	"net/http"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "financial-registry version %s", Version)
}
