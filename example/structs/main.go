package main

import (
	"fmt"
)

type contactInfo struct {
	email 	string
	zipCode int
}

type person struct {
	firstName  	string
	lastName   	string
	contactInfo
}

func main() {
	Jim := person{
		firstName: "Jim",
		lastName: "Hendrix",
		contactInfo: contactInfo{
			email: "jim@email.com",
			zipCode: 10009,
		},
	}

	// &: Give me the memory address of the variable
	// *: Give me the value of address
	JimPoint := &Jim
	JimPoint.updateName("Wut")
	Jim.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}