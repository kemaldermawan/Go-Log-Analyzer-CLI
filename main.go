package main

import (
	"fmt"

	"log-analyzer-go/analyzer"
)

func main() {
	stats, err := analyzer.ParseLogFile("sample.log")
	if err != nil {
		fmt.Printf("Failed to parse log file: %v\n", err)
		return
	}

	stats.PrintReport("0.001s")
}