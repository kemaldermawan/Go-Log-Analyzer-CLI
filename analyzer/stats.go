package analyzer

import "fmt"

type LogEntry struct {
	IP         string
	Method     string
	Endpoint   string
	StatusCode string
}

type LogStats struct {
	TotalLines  int            `json:"total_lines"`
	StatusCodes map[string]int `json:"status_codes"`
	IPRequests  map[string]int `json:"ip_requests"`
	Endpoints   map[string]int `json:"endpoints"`
}

func NewLogStats() *LogStats {
	return &LogStats{
		StatusCodes: make(map[string]int),
		IPRequests:  make(map[string]int),
		Endpoints:   make(map[string]int),
	}
}

func (s *LogStats) AddEntry(entry LogEntry) {
	s.TotalLines++
	s.StatusCodes[entry.StatusCode]++
	s.IPRequests[entry.IP]++
	s.Endpoints[entry.Endpoint]++
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

	fmt.Println("\n[+] IP Requesters:")
	for ip, count := range s.IPRequests {
		fmt.Printf("    - IP %s : %d requests\n", ip, count)
	}

	fmt.Println("\n[+] Requested Endpoints:")
	for ep, count := range s.Endpoints {
		fmt.Printf("    - %s : %d hits\n", ep, count)
	}
	fmt.Println("==================================================")
}