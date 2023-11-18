package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p *person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(firstName string) {
	p.firstName = firstName
}

func main() {
	jimPtr := person{
		firstName:   "Jim",
		lastName:    "Essiet",
		contactInfo: contactInfo{email: "jim@gmail.com", zipCode: 00234},
	}

	jimPtr.updateFirstName("Jimmy")
	jimPtr.print()
}
