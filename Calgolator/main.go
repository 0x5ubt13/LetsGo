package main

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the first number: ")
	aInput, _ := reader.ReadString('\n')
	a, err := strconv.ParseFloat(strings.TrimSpace(aInput), 64)
	if err != nil {
		panic(err)
	} 
	fmt.Print("Enter the second number: ")
	bInput, _ := reader.ReadString('\n')
	b, err := strconv.ParseFloat(strings.TrimSpace(bInput), 64)
	if err != nil {
		panic(err)
	}
	adding(a, b)
}

func adding(a float64, b float64) {
	c := a + b
	fmt.Printf("The sum of %v and %v equals %v\n", a, b, c)
}