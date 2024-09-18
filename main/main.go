package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

const columnSize = 80

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
				if len(records[i][j]) < columnSize {
					preferredLength[j] = len(records[i][j])
				} else {
					preferredLength[j] = columnSize
				}
			}
		}
	}

	for _, record := range records {
		fmt.Println(printDefaultLine(preferredLength))

		maximumLines := findMaximumLine(record)
		if maximumLines == 1 {
			fmt.Println(printLine(record, preferredLength))
		} else {
			printBasedOnMaximumLines(record, preferredLength, maximumLines)
		}
	}
	fmt.Println(printDefaultLine(preferredLength))

}

func printBasedOnMaximumLines(record []string, preferredLength []int, maximumLines int) {
	for i := 0; i < maximumLines; i++ {
		//fmt.Println("maxim", maximumLines)
		line := "|"
		for z, r := range record {
			arr := wrapTextToArray(r, columnSize)
			startLine := (maximumLines - len(arr)) / 2
			arr = shiftArray(arr, startLine)

			//fmt.Println("arr", len(arr), arr, i)
			if len(arr)-1 >= i {
				line += arr[i]
				if len(arr[i]) < preferredLength[z] {
					spacesNumber := preferredLength[z] - len(arr[i])
					line += strings.Repeat(" ", spacesNumber)
				}
			} else {
				line += strings.Repeat(" ", preferredLength[z])
			}

			line += "|"
		}

		fmt.Println(line)
	}
}

func shiftArray(arr []string, startLine int) []string {
	// Create a new array with the same length as the original array plus the startLine offset
	newArr := make([]string, len(arr)+startLine)

	// Fill the new array with empty strings before the startLine
	for i := 0; i < startLine; i++ {
		newArr[i] = ""
	}

	// Copy elements from the original array to the new array starting from startLine
	for i := 0; i < len(arr); i++ {
		newArr[startLine+i] = arr[i]
	}

	return newArr
}

func findMaximumLine(record []string) int {
	maximumLines := 1
	for _, r := range record {
		if len(r) > columnSize {
			arr := wrapTextToArray(r, columnSize)
			if len(arr) > maximumLines {
				maximumLines = len(arr)
			}
		}
	}

	return maximumLines
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

func wrapTextToArray(sentence string, colSize int) []string {
	words := strings.Fields(sentence) // Split the sentence into words
	var result []string
	var line string

	for _, word := range words {
		// If adding the next word exceeds the column size, add the current line to result and reset it
		if len(line)+len(word)+1 > colSize {
			result = append(result, strings.TrimSpace(line)) // Append the current line
			line = word                                      // Start new line with the current word
		} else {
			if len(line) > 0 {
				line += " " // Add space between words
			}
			line += word // Append the word to the current line
		}
	}

	// Append the last line if any
	if len(line) > 0 {
		result = append(result, line)
	}

	return result
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
