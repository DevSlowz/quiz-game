package quiz_engine

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func Read() {

	// Create map to hold test questions and answer
	testQuestions := make(map[string]string)

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
		parts := strings.Split(record[0], ",")
		question := strings.TrimSpace(parts[0])
		answer := strings.TrimSpace(parts[1])
		// fmt.Printf("The part is : %s\n", parts)
		// fmt.Printf("The question is %s the answer is %s\n", question, answer)
		testQuestions[question] = answer
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

// func sortTest([][]string) map[string]string {

// }
