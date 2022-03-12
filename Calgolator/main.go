package main

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"os"
	"math"
)

func main() {
	operators, operation := readingInput()
	total := execute(operators, operation)

	fmt.Printf("The total is %v", total)
}

func readingInput() ([]float64, string) {
	// Making slice with operators, max 100
	operators := make([]float64, 1, 100)

	// Creating the reader buffer so we can use it afterwards 
	reader := bufio.NewReader(os.Stdin)

	// Read first number
	fmt.Print("Enter first number: ")
	firstNumber, _ := reader.ReadString('\n')
	fN, err := strconv.ParseFloat(strings.TrimSpace(firstNumber), 64)
	if err != nil {
		panic(err)
	}

	// Add the new number to the slice
	operators = append(operators, fN)

	for {
		// Read new number
		fmt.Print("Enter new number: ")
		newNumber, _ := reader.ReadString('\n')
		nN, err := strconv.ParseFloat(strings.TrimSpace(newNumber), 64)
		if err != nil {
			panic(err)
		}

		// Add the new number to the slice
		operators = append(operators, nN)

		// Control flow check 
		fmt.Print("Enter more numbers? y/n: ")
		option, err := reader.ReadString('\n')
		choice := strings.ToLower(strings.TrimSpace(option))
		if err != nil {
			panic(err)
		} else if choice == "n" {
			// Select the operation and exit the loop
			fmt.Printf("[+] Numbers correctly processed.\n[+] Using the following as operators:\n%+v\n", operators[1:])
			fmt.Print("Now please select the operation you want to perform (+, -, *, /): ")
			sign, _ := reader.ReadString('\n')
			operation := strings.TrimSpace(sign)
			return operators, operation
		} else {
			continue
		}
	}
}

func execute(operators []float64, operation string) (float64) {
	// First number in slice is "0" so to avoid math problems we need to start counting from index 1 onwards
	firstOperator := operators[1]
	otherOperators := operators[2:]
	for _, operator := range otherOperators {
		switch operation {
		case "+":
			firstOperator += operator
		case "-":
			firstOperator -= operator
		case "*":
			firstOperator *= operator
		case "/":
			firstOperator /= operator
		} 
	}
	
	return (math.Round(firstOperator*100) / 100)
}