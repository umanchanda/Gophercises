package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Quiz is a quiz with a simple math question followed by a comma and then an answer
type Quiz struct {
	Question string
	Answer   string
}

func main() {
	problems, err := readCSV("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	var correct = 0
	quizproblems := parseSomeShitIDK(problems)
	for _, problem := range quizproblems {
		fmt.Print(problem.Question + " = ")
		var guess = ""
		fmt.Scanf("%v", &guess)
		if guess == problem.Answer {
			correct++
		}
	}
	fmt.Printf("You got: %d right!\n", correct)
}

func readCSV(filename string) ([][]string, error) {
	csvfile, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer csvfile.Close()

	lines, err := csv.NewReader(csvfile).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func parseSomeShitIDK(problems [][]string) []Quiz {
	quizproblems := make([]Quiz, len(problems))
	for a, b := range problems {
		quizproblems[a] = Quiz{
			Question: b[0],
			Answer:   b[1],
		}
	}
	return quizproblems
}
