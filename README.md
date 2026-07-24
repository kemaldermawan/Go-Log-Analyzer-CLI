# High-Performance Log Analyzer CLI (Go)

A fast, memory-efficient Command Line Interface (CLI) tool built in **Go (Golang)** for parsing, analyzing, and detecting security anomalies in server log files (Nginx/Apache format).

This tool is engineered using **buffered streaming I/O** to process large log files (hundreds of megabytes) with minimal RAM consumption.

## Key Features

* **Memory-Efficient Parsing**: Reads log files line-by-line using `bufio.Scanner` to maintain a low memory footprint regardless of file size.
* **Pattern Extraction**: Uses Regular Expressions (Regex) to extract IP addresses, HTTP methods, endpoints, and status codes.
* **Top-N Analytics**: Sorts and identifies the **Top 5 IP Requesters** and **Top 5 Most Visited Endpoints**.
* **Security Anomaly Detection**: Automatically flags IP addresses exhibiting suspicious behavior (e.g., high volume of `401 Unauthorized` or `404 Not Found` responses indicating potential brute force or directory scanning).
* **Multi-Format Export**: Supports exporting analytical reports directly into **JSON** or **CSV** formats.

## Tech Stack & Architecture

* **Language**: Go (Golang 1.20+)
* **Standard Libraries**: `os`, `bufio`, `regexp`, `flag`, `encoding/json`, `encoding/csv`, `sort`, `time`
* **Architecture**: Modular layout dividing CLI routing, parsing logic, analytical aggregation, and report exporting.

## Roadmap & Development Phases

To ensure clear and structured implementation, this project is built across **6 core phases**:

### Phase 1: Environment Setup & Module Initialization
Initialize the Go module (`go mod init log-analyzer-go`) and configure the project directory structure (`main.go`, `sample.log`, and the `analyzer/` package).

### Phase 2: Log File Parsing (Streaming File I/O & Regex)
Implement `ParseLogFile` in `analyzer/parser.go` using `bufio.Scanner` to read log lines sequentially. Use Regular Expressions to extract IP, HTTP method, endpoint, and status code from each line.

### Phase 3: Data Aggregation & Top-N Analytics
Build `analyzer/stats.go` using Go `map` types to aggregate total lines, status codes, IP request counts, and endpoint hits. Implement slice sorting (`sort.Slice`) to extract Top 5 rankings.

### Phase 4: Security Anomaly & Brute Force Detection
Construct `analyzer/security.go` to inspect error rates per IP. Flag addresses exceeding a defined threshold of `401 Unauthorized` or `404 Not Found` responses as potential scanners or brute-force attackers.

### Phase 5: Multi-Format Exporter & CLI Flags
Implement `analyzer/exporter.go` to write analytical results into structured `JSON` or `CSV` files. Configure `main.go` using the `flag` package to parse user inputs (`-file`, `-export`, `-output`).

### Phase 6: Error Handling, Benchmarking & Documentation
Add robust error handling for missing files or invalid inputs. Measure execution duration (`time.Since`) and document system capabilities within `README.md`.

## Installation & Usage

### Prerequisites
Ensure you have **Go 1.20+** installed on your system.

### 1. Clone the Repository
```bash
git clone [https://github.com/kemaldermawan/Go-Log-Analyzer-CLI.git](https://github.com/kemaldermawan/Go-Log-Analyzer-CLI.git)
cd Go-Log-Analyzer-CLI
```

### 2. Run with Go CLI
To analyze a log file:
```bash
go run main.go -file=sample.log
```

To analyze and export the report to a JSON or CSV file:
```bash
# Export to JSON
go run main.go -file=sample.log -export=json -output=report.json

# Export to CSV
go run main.go -file=sample.log -export=csv -output=report.csv
```

### 3. Build Executable Binary
Compile a standalone binary executable:
```bash
go build -o log-analyzer main.go
./log-analyzer -file=sample.log
```

## Sample CLI Output

```text
==================================================
           LOG ANALYZER REPORT (GO)               
==================================================
Total Lines Processed : 15,420 lines
Execution Time        : 0.08 seconds

[+] HTTP Status Codes Breakdown:
    - Status 200 OK          : 12,300
    - Status 404 Not Found   : 2,800
    - Status 401 Unauthorized: 200
    - Status 500 Error       : 120

[+] Top 5 IP Requesters:
    1. 192.168.1.45   - 3,420 requests
    2. 10.0.0.12      - 2,100 requests
    3. 172.16.0.8     - 1,850 requests
    4. 192.168.1.10   - 1,200 requests
    5. 10.0.0.5       - 890 requests

[+] Top 5 Requested Endpoints:
    1. GET  /api/v1/products   - 4,500 hits
    2. POST /api/v1/login      - 3,100 hits
    3. GET  /checkout          - 2,100 hits
    4. GET  /dashboard         - 1,200 hits
    5. GET  /admin             - 800 hits

[!] SECURITY ALERTS / ANOMALIES DETECTED:
    [ALERT] IP 10.0.0.12 flagged for Potential Brute Force / Scanning (150 Failed Attempts).

[✓] Report successfully exported to: report.json
==================================================
```
