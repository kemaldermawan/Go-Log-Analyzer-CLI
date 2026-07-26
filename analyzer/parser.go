package analyzer

import (
	"bufio"
	"os"
	"regexp"
)

var logPattern = regexp.MustCompile(`^(\S+) \S+ \S+ \[.*?\] "(\S+) (\S+) \S+" (\d{3}) \d+`)

func ParseLogFile(filePath string) (*LogStats, error) {
	// 1. Buka file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := NewLogStats()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		matches := logPattern.FindStringSubmatch(line)

		if len(matches) == 5 {
			entry := LogEntry{
				IP:         matches[1],
				Method:     matches[2],
				Endpoint:   matches[3],
				StatusCode: matches[4],
			}

			stats.AddEntry(entry)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
