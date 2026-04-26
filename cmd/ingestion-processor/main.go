package main

import (
	"fmt"
	"os"

)

func main() {
	fmt.Println("Ingestion Processor Started")

	if len(os.Args) <2 {
		fmt.Println("No input file provided")
		return
	}

	input := os.Args[1]

	fmt.Println("Processing file:", input)
}

