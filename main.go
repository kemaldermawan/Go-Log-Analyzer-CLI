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
	exportPath := flag.String("export", "", "Path to export the report as JSON (optional)")

	flag.Parse()

	start := time.Now()

	stats, err := analyzer.ParseLogFile(*logFile)
	if err != nil {
		fmt.Printf("Error: Failed to parse log file '%s': %v\n", *logFile, err)
		return
	}

	duration := time.Since(start).String()

	stats.PrintReport(duration, *threshold)

	if *exportPath != "" {
		err := stats.ExportToJSON(*exportPath)
		if err != nil {
			fmt.Printf("Error exporting report: %v\n", err)
		}
	}
}
