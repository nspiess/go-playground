package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	p := person{
		firstName: "Max",
		lastName:  "Muster",
		contactInfo: contactInfo{
			email:   "max@muster.com",
			zipCode: 1111,
		},
	}

	p.updateFirstName("alex")
	p.print()
}

func (p *person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(firstName string) {
	// old go? (*p) // -> unpack pointer to a value
	p.firstName = firstName
}
