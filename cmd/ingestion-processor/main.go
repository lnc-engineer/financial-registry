package main

import (
	"fmt"
	"os"
	"strings"

)

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

		//split into lines
		lines :=strings.Split(content, "\n")

		for _, line := range lines {
			if line == "" {
				continue
			}

			fmt.Println("Record:", line)
		}
	}

}