package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}
type person struct {
	firstName string
	lastName string
    contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// Shortcut in Go that allows a method with a receiver of a pointer of a type
	// to be called with just the type
	jim.updateName("Jimmy")
	jim.print()
}

// Update first name of person using pointers
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Print out a person
func (p person) print() {
	fmt.Printf("%+v", p)
}