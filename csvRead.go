package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file
	file, err := os.Open("Cricket_Players_Stats.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the rows in the CSV file
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print all the rows
	fmt.Println("Printing all rows:")
	for _, row := range rows {
		fmt.Println(row)
	}

	// Print all the rows except the header row
	fmt.Println("\nPrinting all rows except the header row:")
	for i, row := range rows {
		if i == 0 {
			continue
		}
		fmt.Println(row)
	}
}
