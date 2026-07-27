package analyzer

import (
	"encoding/json"
	"fmt"
	"os"
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

	fmt.Printf("[✓] Report successfully exported to: %s\n", outputPath)
	return nil
}
