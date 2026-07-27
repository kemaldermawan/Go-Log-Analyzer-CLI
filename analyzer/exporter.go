package analyzer

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func (s *LogStats) ExportToJSON(outputPath string) error {
	jsonData, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}

	err = os.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	fmt.Printf("[✓] JSON report successfully exported to: %s\n", outputPath)
	return nil
}

func (s *LogStats) ExportToCSV(outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Category", "Key", "Value"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	writer.Write([]string{"General", "Total Lines", strconv.Itoa(s.TotalLines)})

	for k, v := range s.StatusCodes {
		writer.Write([]string{"Status Code", k, strconv.Itoa(v)})
	}

	for k, v := range s.IPRequests {
		writer.Write([]string{"IP Request", k, strconv.Itoa(v)})
	}

	for k, v := range s.Endpoints {
		writer.Write([]string{"Endpoint", k, strconv.Itoa(v)})
	}

	for k, v := range s.FailedRequests {
		writer.Write([]string{"Security Alert (Failed IP)", k, strconv.Itoa(v)})
	}

	fmt.Printf("[✓] CSV report successfully exported to: %s\n", outputPath)
	return nil
}
