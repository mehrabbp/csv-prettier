package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filePath := os.Args[1]

	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to open the file: %v", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read the CSV file: %v", err)
	}

	preferredLength := make([]int, len(records[0]))
	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[i]); j++ {
			if len(records[i][j]) > preferredLength[j] {
				preferredLength[j] = len(records[i][j])
			}
		}
	}

	for _, record := range records {
		fmt.Println(printDefaultLine(preferredLength))
		fmt.Println(printLine(record, preferredLength))
	}
	fmt.Println(printDefaultLine(preferredLength))

}

func printLine(record []string, preferredLength []int) string {
	line := "|"
	for i, r := range record {
		line += r
		spacesNumber := preferredLength[i] - len(r)
		line += strings.Repeat(" ", spacesNumber)
		line += "|"
	}

	return line
}

func printDefaultLine(length []int) string {
	line := ""
	line += "+"
	minesNumbers := 0
	for _, l := range length {
		minesNumbers += l
	}
	minesNumbers += len(length)
	minesString := "-"
	line += strings.Repeat(minesString, minesNumbers)
	line += "+"

	return line
}
