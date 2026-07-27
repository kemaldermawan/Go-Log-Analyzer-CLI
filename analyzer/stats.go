package analyzer

import (
	"fmt"
	"sort"
)

type LogEntry struct {
	IP         string
	Method     string
	Endpoint   string
	StatusCode string
}

type LogStats struct {
	TotalLines     int            `json:"total_lines"`
	StatusCodes    map[string]int `json:"status_codes"`
	IPRequests     map[string]int `json:"ip_requests"`
	Endpoints      map[string]int `json:"endpoints"`
	FailedRequests map[string]int `json:"failed_requests"`
}

func NewLogStats() *LogStats {
	return &LogStats{
		StatusCodes:    make(map[string]int),
		IPRequests:     make(map[string]int),
		Endpoints:      make(map[string]int),
		FailedRequests: make(map[string]int),
	}
}

func (s *LogStats) AddEntry(entry LogEntry) {
	s.TotalLines++
	s.StatusCodes[entry.StatusCode]++
	s.IPRequests[entry.IP]++
	s.Endpoints[entry.Endpoint]++

	if entry.StatusCode == "401" || entry.StatusCode == "404" {
		s.FailedRequests[entry.IP]++
	}
}

type kv struct {
	Key   string
	Value int
}

func getTopN(m map[string]int, n int) []kv {
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	if len(ss) > n {
		return ss[:n]
	}
	return ss
}

func (s *LogStats) PrintReport(duration string) {
	fmt.Println("==================================================")
	fmt.Println("           LOG ANALYZER REPORT (GO)               ")
	fmt.Println("==================================================")
	fmt.Printf("Total Lines Processed : %d lines\n", s.TotalLines)
	fmt.Printf("Execution Time        : %s\n\n", duration)

	fmt.Println("[+] HTTP Status Codes:")
	for code, count := range s.StatusCodes {
		fmt.Printf("    - Status %s : %d times\n", code, count)
	}

	fmt.Println("\n[+] Top 5 IP Requesters:")
	for _, kv := range getTopN(s.IPRequests, 5) {
		fmt.Printf("    - IP %s : %d requests\n", kv.Key, kv.Value)
	}

	fmt.Println("\n[+] Top 5 Requested Endpoints:")
	for _, kv := range getTopN(s.Endpoints, 5) {
		fmt.Printf("    - %s : %d hits\n", kv.Key, kv.Value)
	}
	fmt.Println("\n[!] SECURITY ALERTS / ANOMALIES DETECTED:")
	alerts := DetectAnomalies(s.FailedRequests, 3)
	if len(alerts) == 0 {
		fmt.Println("    [✓] No anomalies detected.")
	} else {
		for _, alert := range alerts {
			fmt.Printf("    %s\n", alert)
		}
	}
	fmt.Println("==================================================")
}
