package analyzer

import "fmt"

func DetectAnomalies(failedRequests map[string]int, threshold int) []string {
	var alerts []string

	for ip, count := range failedRequests {
		if count >= threshold {
			alert := fmt.Sprintf("[ALERT] IP %s flagged for Potential Brute Force / Scanning (%d Failed Attempts).", ip, count)
			alerts = append(alerts, alert)
		}
	}

	return alerts
}
