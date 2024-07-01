package main

import (
	"github.com/DevSlowz/csv"
)

func main() {

	// Read quiz via CSV file
	test := csv.Read("filename")
	println(test)
}
