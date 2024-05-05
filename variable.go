package main

import "fmt"

func main() {
	var name string // type is required if no value when declaration

	name = "Ridwan Nurul"
	fmt.Println(name)

	name = "Ridwan Ikhsan"
	fmt.Println(name)

	otherName := "Frank" // := same as var, type is optional
	fmt.Println(otherName)

	otherName = "Martin"
	fmt.Println(otherName)

	// multiple var declaration
	var (
		firstName = "Ridwan"
		middleName = "Nurul"
		lastName = "Ikhsan"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
