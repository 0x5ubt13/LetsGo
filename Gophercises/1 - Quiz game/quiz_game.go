package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

// Predeclaring exit triggers
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a .CSV file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	f, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
		os.Exit(1)
	}
	defer f.Close()

	// read the input file
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided .CSV file.")
	}

	problems := parseLines(lines)
	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// Main quiz control flow
	for i, problemo := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problemo.question)

		// Creating an anonymous func and a channel to handle the user input
		answerCh := make(chan string)
		go func() {
			var answr string
			fmt.Scanf("%s\n", &answr)
			answerCh <- answr
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nTime's out! You scored %d out of %d.\n", correct, len(problems))
			return
		case answr := <-answerCh:
			if answr == problemo.answer {
				fmt.Println("Correct!")
				correct++
			}
		}
	}
}

// Parse the input file
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}
