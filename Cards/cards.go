package main

import "fmt"

func main() {
	cards := deck{newCard()}
	fmt.Println(cards)
}

func newCard() string {
	return "Ace of Hearts"
}