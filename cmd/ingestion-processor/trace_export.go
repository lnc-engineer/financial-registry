package main

import (
	"encoding/json"
	"fmt"

	"github.com/lnc-engineer/financial-registry/internal/execution"
)

func PrintTraceJSON(root execution.TraceNode) {

	data, err := json.MarshalIndent(root.Export(), "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
