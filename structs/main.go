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

//Just like C
//&jimPointer gives you the memory address of Jim
//&jimpointer gives you the value of that memory address.

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// jim.print()

	jim.updateName("Jimmy") //This also works, go will auto convert it for you.
	jim.print()
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName //Turning pointerToPerson into an actual object resides in that memory.
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
