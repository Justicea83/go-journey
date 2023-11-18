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

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func main() {
	jim := person{
		firstName:   "Jim",
		lastName:    "Essiet",
		contactInfo: contactInfo{email: "jim@gmail.com", zipCode: 00234},
	}

	jimPointer := &jim
	jimPointer.updateFirstName("Jimmy")
	jim.print()
}
