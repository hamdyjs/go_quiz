package main

import (
	"encoding/csv"
	"log"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {
	// Read from CSV file
	file, _ := os.Open("problems.csv")
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Parse CSV file into problems
	problems := make([]problem, 0)
	for _, record := range records {
		problems = append(problems, problem{record[0], record[1]})
	}
}
