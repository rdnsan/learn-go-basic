package main

import "fmt"

func greeting(name string) string {
	return "Hello, " + name
}

// return multiple value
func getFullName() (string, string) {
	return "Ridwan", "Ikhsan"
}

func main() {
	greet := greeting("Ridwan")
	fmt.Println(greet)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// ignore some return value
	onlyFirstName, _ := getFullName()
	fmt.Println(onlyFirstName)
}
