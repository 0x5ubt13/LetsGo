package main

import (
	"fmt"
)

// Create a function that returns an additional function that returns successive Fibonacci numbers
// Mathematically: Xn = Xn-1 + Xn-2
// Pseudocode: declare two variables inside a function, that returns another function, that returns the sum of the    
func gobonacci() func() int {
	first, second := 0, 1
	return func() int {
		first, second = second, first+second
		return first
	}
}

func main() {
	f := gobonacci()
	for i := 0; i < 92; i++ {
		fmt.Println(f())
	}
}