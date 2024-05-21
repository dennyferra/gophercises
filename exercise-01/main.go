package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := flag.String("file", "problems.csv", "a csv file in the format of 'question,answer'")
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

	for _, record := range records {
		fmt.Printf("What is %s? ", record[0])

		var answer string
		n, err := fmt.Scanf("%s\n", &answer)
		if err != nil || n != 1 {
			fmt.Println("Failed to read the answer")
		}

		if answer == record[1] {
			correct++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}
	}

	fmt.Println("\nYou answered", correct, "out of", len(records), "questions correctly.")
}
