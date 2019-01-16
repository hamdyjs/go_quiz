package main

import (
	"encoding/csv"
	"fmt"
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

	fmt.Println("=== STARTING QUIZ ===")
	questionCounter := 1
	correctAnswers := 0
	for _, problem := range problems {
		fmt.Printf("Question #%d: %s\n", questionCounter, problem.question)
		var answer string
		fmt.Scan(&answer)
		if answer == problem.answer {
			correctAnswers++
		}
		questionCounter++
	}

	fmt.Println("=== FINISHED QUIZ ===")
	fmt.Printf("Result: %d/%d", correctAnswers, questionCounter-1)
}
