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
}

func main() {
	// alex := person{"Alex", "Anderson"}	// not recommended, as we do may not know the order of the attributes of struc
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	// fmt.Printf("%+v", alex)
}

// func main() {
// 	jim := person{
// 		firstName: "Jim",
// 		lastName:  "Party",
// 		contact: contactInfo{
// 			email:   "jim@gmail.com",
// 			zipCode: 94000,
// 		},
// 	}

// 	jimPoiner := &jim

// 	jimPoiner.updateName("jimmy")
// 	jim.print()
// }

// func (pointerToPerson *person) updateName(newFirstName string) {
// 	(*pointerToPerson).firstName = newFirstName
// }

// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }
