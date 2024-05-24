package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	filePath := flag.String("file", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*filePath)

	if err != nil {
		fmt.Println("Failed to open the CSV file: ", *filePath)
		os.Exit(1)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Failed to parse the provided CSV file")
		os.Exit(1)
	}

	correct := 0

	fmt.Println("Press enter to start the quiz")
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, record := range records {
		fmt.Printf("Question %d: %s = ", i+1, record[0])
		answerCh := make(chan string)
		go func() {
			var answer string
			n, err := fmt.Scanf("%s\n", &answer)
			if err != nil || n != 1 {
				fmt.Println("Failed to read the answer")
			}
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nYou answered", correct, "out of", len(records), "questions correctly.")
			return
		case answer := <-answerCh:
			if answer == record[1] {
				correct++
				fmt.Println("  Correct!")
			} else {
				fmt.Println("  Incorrect!")
			}
		}
	}

}
