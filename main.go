package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	nortonCsv := "yournortonpasswordmanager.csv"
	chromeCsv := "chrome_passwords.csv"

	file, err := os.Open(nortonCsv)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	if _, err := reader.Read(); err != nil {
		log.Fatalf("Failed to read metadata: %v", err)
	}
	if _, err := reader.Read(); err != nil {
		log.Fatalf("Failed to read metadata: %v", err)
	}

	outputFile, err := os.Create(chromeCsv)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	chromeHeader := []string{"name", "url", "username", "password", "note"}
	if err := writer.Write(chromeHeader); err != nil {
		log.Fatalf("Failed to write header: %v", err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if len(record) < 6 {
			log.Printf("Skipping malformed record: %v\n", record)
			continue
		}

		name := record[3]
		url := record[4]
		username := record[1]
		password := record[2]
		note := record[5]

		chromeRecord := []string{name, url, username, password, note}
		if err := writer.Write(chromeRecord); err != nil {
			log.Fatalf("Failed to write record: %v", err)
		}
	}

	fmt.Printf("Conversion complete. File saved as %s\n", chromeCsv)
}
