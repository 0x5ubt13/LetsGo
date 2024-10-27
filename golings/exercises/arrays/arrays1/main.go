// arrays1
// Make me compile!

package main

import "fmt"

func main() {
	var colours [3]string

	colours[0] = "red"
	colours[1] = "green"
	colours[2] = "blue"

	fmt.Printf("First colour is %s\n", colours[0])
	fmt.Printf("Last colour is %s\n", colours[len(colours)-1])
}
