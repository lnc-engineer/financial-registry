package main

import (
	"fmt"
	"os"
	"strings"
)

type Record struct {
	Name string
	Age  string
	Role string
}

func main() {
	fmt.Println("Ingestion Processor Started")

	if len(os.Args) < 2 {
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

		content := string(data)

		// Split into lines
		lines := strings.Split(content, "\n")

		for _, line := range lines {
			if line == "" {
				continue
			}

			fields := strings.Split(line, ",")

			if len(fields) != 3 {
				fmt.Println("Skipping invalid record:", line)
				continue
			}

			record := Record{
				Name: strings.TrimSpace(fields[0]),
				Age:  strings.TrimSpace(fields[1]),
				Role: strings.TrimSpace(fields[2]),
			}

			fmt.Printf("Parsed Record: %+v\n", record)
		}
	}
}
