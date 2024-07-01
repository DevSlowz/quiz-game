package main

import (
	"flag"

	"github.com/DevSlowz/quiz_engine"
)

func main() {
	flags := flag.String("file", "problems.csv", "Path to CSV containing test questions")
	flag.Parse()
	// Read quiz via CSV file
	// test := csv.Read("filename")
	// println(test)

	// Give user test

	// Start
	quiz_engine.StartQuiz(quiz_engine.Read(*flags))

	// Get use test results
	// quiz_engine.Results()
}
