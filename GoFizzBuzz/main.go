package main

import "fmt"

func main() {
	for i := 1; i < 20; i++ {
		if i % 3 == 0 {
			if i % 5 == 0 {
				fmt.Println("Fizzbuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}