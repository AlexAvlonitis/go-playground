package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "alex",
		lastName:  "xela",
		contactInfo: contactInfo{
			email: "asd@asd.com",
			zip:   123,
		},
	}

	alex.updateName("nick")
	alex.print()
}

func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).firstName = name
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v\n", pointerToPerson)
}
