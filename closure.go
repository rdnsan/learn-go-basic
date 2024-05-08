package main

import "fmt"

func main() {
	id := 1

	autoIncrementId := func() {
		fmt.Println("increment")
		id++
	}

	fmt.Println(id)

	autoIncrementId()
	autoIncrementId()
	autoIncrementId()

	fmt.Println(id)
}
