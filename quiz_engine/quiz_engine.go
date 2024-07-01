package quiz_engine

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Read() {
	// Open specified file
	file, err := os.Open("./problems.csv")

	// if the file is not found or cant open log details
	if err != nil {
		log.Fatal(err)
	}

	// Executed at the end of Read function
	defer closeFile(file)

	// parse csv
	reader := csv.NewReader(file)

	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}

// Close file and verify it closed correctly
func closeFile(f *os.File) {
	fmt.Println("Closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}