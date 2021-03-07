package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo // field name contactInfo with type contactInfo
}

func main() {

	/*
		// three ways to declare a struct
		// alex := person{"Alex", "Anderson"}
		// alex := person{firstName: "Alex", lastName: "Anderson"}
		var alex person
		alex.firstName = "alex"
		alex.lastName = "anderson"
		fmt.Println(alex)
		// %+v will print out all the different field name and their values
		fmt.Printf("%+v", alex)
	*/

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// contact: contactInfo{
		// 	email:   "jim@gmail.com",
		// 	zipCode: 94000,
		// },
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	// jim.print()

	// shortcut in pointers in go
	jim.updateName("jimmy")
	jim.print()

	/*
		Whenever you pass an integer, float, string, or struct into a function, what does Go do with that argument?
		it creates a copy of each argument, and those copies are used inside of the function

		value types { int, float, string, bool, structs }
		reference types { slices, maps, channels, pointers, function }

		even though reference value is copied, it still points to the location where the value is in RAM
	*/
	//big gocha with go
	mySlice := []string{"Hi", "There", "How", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// go is a pass by value language whenever we pass some value, a copy will be create in RAM
func (pointerToPerson *person) updateName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}
