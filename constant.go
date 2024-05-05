package main

import "fmt"

func main()  {
	const greeting string = "Hello,"

	const (
		firstName = "Ridwan"
		lastName = "Ikhsan"
	)

	fmt.Println(greeting, firstName + " " + lastName)
}
