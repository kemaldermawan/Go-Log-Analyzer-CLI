package main

import (
	"flag"
	"fmt"
	"time"

	"log-analyzer-go/analyzer"
)

func main() {
	logFile := flag.String("file", "sample.log", "Path to the log file to analyze")
	threshold := flag.Int("threshold", 3, "Failed request threshold for anomaly detection")
	jsonPath := flag.String("json", "", "Path to export the report as JSON (optional)")
	csvPath := flag.String("csv", "", "Path to export the report as CSV (optional)")

	flag.Parse()
	start := time.Now()

	stats, err := analyzer.ParseLogFile(*logFile)
	if err != nil {
		fmt.Printf("Error: Failed to parse log file '%s': %v\n", *logFile, err)
		return
	}

	duration := time.Since(start).String()
	stats.PrintReport(duration, *threshold)

	if *jsonPath != "" {
		err := stats.ExportToJSON(*jsonPath)
		if err != nil {
			fmt.Printf("Error exporting JSON report: %v\n", err)
		}
	}

	if *csvPath != "" {
		err := stats.ExportToCSV(*csvPath)
		if err != nil {
			fmt.Printf("Error exporting CSV report: %v\n", err)
		}
	}
}
