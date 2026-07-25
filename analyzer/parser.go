package analyzer

import (
	"bufio"
	"os"
	"regexp"
)

// Pola Regex untuk mengambil: [1] IP, [2] Method, [3] Endpoint, [4] Status Code
var logPattern = regexp.MustCompile(`^(\S+) \S+ \S+ \[.*?\] "(\S+) (\S+) \S+" (\d{3}) \d+`)

// ParseLogFile membaca berkas log baris demi baris dan mengembalikan hasil hitungan statistik
func ParseLogFile(filePath string) (*LogStats, error) {
	// 1. Buka file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Pastikan file ditutup kembali setelah fungsi selesai

	stats := NewLogStats()
	scanner := bufio.NewScanner(file)

	// 2. Baca file baris demi baris
	for scanner.Scan() {
		line := scanner.Text()

		// 3. Cocokkan baris dengan pola Regex
		matches := logPattern.FindStringSubmatch(line)

		// Jika pola cocok (menghasilkan 5 elemen: string utuh + 4 grup tangkapan)
		if len(matches) == 5 {
			entry := LogEntry{
				IP:         matches[1],
				Method:     matches[2],
				Endpoint:   matches[3],
				StatusCode: matches[4],
			}

			// 4. Tambahkan ke perhitungan statistik
			stats.AddEntry(entry)
		}
	}

	// Cek apakah ada error saat proses pemindaian baris
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}