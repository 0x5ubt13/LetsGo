package main

import("fmt")

func main() {
	gab := person{firstName: "Gab", lastName: "Gar"}
	fmt.Println(gab)
}

type person struct {
	firstName string
	lastName string
}
