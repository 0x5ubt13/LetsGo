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
	// Creating the reader buffer so we can use it afterwards 
	reader := bufio.NewReader(os.Stdin)

	// First number
	fmt.Print("Enter the first number: ")
	aInput, _ := reader.ReadString('\n')
	a, err := strconv.ParseFloat(strings.TrimSpace(aInput), 64)
	if err != nil {
		panic(err)
	} 

	// Second number
	fmt.Print("Enter the second number: ")
	bInput, _ := reader.ReadString('\n')
	b, err := strconv.ParseFloat(strings.TrimSpace(bInput), 64)
	if err != nil {
		panic(err)
	}

	// Now that the numbers are parsed, add them up calling adding()
	adding(a, b)
}

func adding(a float64, b float64) {
	// Sum the values passed 
	c := a + b
	
	// Making sure we have a precise decimal number
	d := math.Round(c*100) / 100

	fmt.Printf("The sum of %v and %v equals %v\n", a, b, d)
}