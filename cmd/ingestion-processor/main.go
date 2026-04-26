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

	for i := 1; i < len(os.Args); i++ {
		file := os.Args[i]

		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		fmt.Println("Processing file:", file)
		fmt.Println(string(data))
	}

}