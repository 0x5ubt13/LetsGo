package main

import "fmt"

// Upgrade the bots with getGreeting() to bot type
type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// Allow all bots to use printGreeting
func printGreeting(b bot) string {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}