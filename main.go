package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("file", "problems.csv", "Problems CSV file")
	answerLimit := flag.Int("time", 30, "Time limit to answer each question")
	shuffleProblems := flag.Bool("shuffle", false, "Shuffle problems?")
	flag.Parse()

	// Read from CSV file
	file, _ := os.Open(*fileName)
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
	if *shuffleProblems {
		rand.Seed(time.Now().UTC().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	fmt.Println("=== STARTING QUIZ ===")
	questionCounter := 1
	correctAnswers := 0
quizLoop:
	for _, problem := range problems {
		answerTimer := time.NewTimer(time.Duration(*answerLimit) * time.Second)
		scanChan := make(chan struct{})

		fmt.Printf("Question #%d: %s\n", questionCounter, problem.question)
		var answer string

		go func() {
			fmt.Scan(&answer)
			scanChan <- struct{}{}
		}()

		select {
		case <-answerTimer.C:
			break quizLoop
		case <-scanChan:
			answerTimer.Stop()
			if answer == problem.answer {
				correctAnswers++
			}
			questionCounter++
		}
	}

	fmt.Println("=== FINISHED QUIZ ===")
	fmt.Printf("Result: %d/%d", correctAnswers, len(problems))
}
