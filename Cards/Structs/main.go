package main

import (
	"fmt"
)

func (p person) print() {
	fmt.Printf("%+v", p)
}

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	hackerAlias string
	contact contactInfo // embedded struct
}

func main() {
	var zc person
	zc.firstName = "Dade"
	zc.lastName = "Murphy"
	zc.hackerAlias = "Zero Cool"
	
	mr := person{firstName: "Elliot", lastName: "Alderson"}
	mr.hackerAlias = "Mr. Robot"

	n := person{
		firstName: "Thomas",
		lastName: "Anderson",
		hackerAlias: "Neo",
		contact: contactInfo{
			email: "thechosenone@matrix.com",
			zipCode: 99999,
		},
	}

	fmt.Printf("%+v\n", zc)
	fmt.Println(mr)
	n.print()
}
