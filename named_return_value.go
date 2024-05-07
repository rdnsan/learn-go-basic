package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Ridwan"
	lastName = "Ikhsan"

	return firstName, middleName, lastName
}

func main() {
	first, middle, _ := getCompleteName()
	fmt.Println(first, middle)
}
