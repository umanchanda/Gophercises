package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	quizPtr := flag.String("csv", "problems.csv", "Quiz filename")
	timePtr := flag.Int("limit", 30, "Quiz time limit")
	flag.Parse()

	numbercorrect := 0
	quiz := parseCSV(*quizPtr)

	timeout := make(chan bool, *timePtr)
	go func() {
		for _, record := range quiz {
			fmt.Printf("%s=", record[0])
			var userinput string
			fmt.Scanln(&userinput)

			if userinput == record[1] {
				numbercorrect += 1
			}
		}
		timeout <- true
	}()

	select {
	case <-timeout:
		fmt.Println("All questions answered")
	case <-time.After(time.Duration(*timePtr) * time.Second):
		fmt.Println("Times up")
	}

	fmt.Printf("You got %d out of %d correct\n", numbercorrect, len(quiz))
}

func parseCSV(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to open file")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	quiz, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Unable to read CSV")
	}
	return quiz
}
