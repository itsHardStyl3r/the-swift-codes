package tools

import (
	"encoding/csv"
	"github.com/charmbracelet/log"
	"os"
)

func ReadCsv(filePath string, skipHeader bool) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to read the file '%v'.\n%v", filePath, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Failed to close the file '%v'.\n%v", filePath, err)
		}
	}(f)
	csvReader := csv.NewReader(f)
	if skipHeader { // Skipping headers.
		_, _ = csvReader.Read()
	}
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse csv file '%v'.\n%v", filePath, err)
	}
	return records
}
