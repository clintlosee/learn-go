package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Johnson"}
	// fmt.Println(alex)

	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Johnsonsons"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "George",
		contact: contactInfo{
			email:   "Jim@email.com",
			zipCode: 94000,
		},
		// contactInfo: contactInfo{
		// 	email:   "Jim@email.com",
		// 	zipCode: 94000,
		// },
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
