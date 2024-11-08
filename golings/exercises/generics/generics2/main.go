// generics2
// Make me compile!

package main

import "fmt"

func main() {
	fmt.Println(addNumbers(1, 2))
	fmt.Println(addNumbers(1.0, 2.3))
}

func addNumbers[T int | int64 | float64](n1, n2 T) any {
	return n1 + n2
}
