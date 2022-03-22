package main

import "fmt"

func main() {
	var emptyMap map[string]int
	fmt.Println(emptyMap)

	anotherEmptyMap := make(map[int]int)
	fmt.Println(anotherEmptyMap)

	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colours["white"] = "#ffffff"
	printMap(colours)

	delete(colours, "green")
	fmt.Println(colours)
}

func printMap(c map[string]string) {
	x := 1
	for colour, hex := range c {
		fmt.Printf("Colour #%v: %v. Hex value: %v\n", x, colour, hex)
		x++
	}
}