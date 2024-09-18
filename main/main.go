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
				preferredLength[j] = len(records[i][j])
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

			//fmt.Println("arr", len(arr), arr, i)
			if len(arr)-1 >= i {
				line += arr[i]
				if len(arr[i]) < preferredLength[z] {
					if preferredLength[z] < columnSize {
						spacesNumber := preferredLength[z] - len(arr[i])
						line += strings.Repeat(" ", spacesNumber)
					} else {
						spacesNumber := columnSize - len(arr[i])
						line += strings.Repeat(" ", spacesNumber)
					}
				}
			} else {
				if preferredLength[z] < columnSize {
					line += strings.Repeat(" ", preferredLength[z])
				} else {
					line += strings.Repeat(" ", columnSize)
				}
			}

			line += "|"
		}

		fmt.Println(line)
	}
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

		if preferredLength[i] < columnSize {
			spacesNumber := preferredLength[i] - len(r)
			line += strings.Repeat(" ", spacesNumber)
		} else {
			spacesNumber := columnSize - len(r)
			line += strings.Repeat(" ", spacesNumber)
		}
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
		if l > columnSize {
			minesNumbers += columnSize
		} else {
			minesNumbers += l
		}
	}
	minesNumbers += len(length)
	minesString := "-"
	line += strings.Repeat(minesString, minesNumbers)
	line += "+"

	return line
}
