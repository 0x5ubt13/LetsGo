package main

import "fmt"

func main() {
	intSlice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, d := range intSlice {
		if d % 2 == 0 {
			fmt.Printf("%d is even\n", d)
		} else {
			fmt.Printf("%d is odd\n", d)
		}
	}
}
