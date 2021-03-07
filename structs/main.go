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
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v", jim)
}
