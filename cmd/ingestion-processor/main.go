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

	fmt.Println("Number of inputs:", len(os.Args)-1)

	for i := 1; i < len(os.Args); i++ {

		fmt.Println("Processing file:", os.Args[i])
	}

}