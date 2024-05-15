package main

import "fmt"

type Car struct {
	Brand, Model string
}

func main() {
	car1 := new(Car)
	car2 := car1

	car2.Model = "GTR"

	fmt.Println(car1)
	fmt.Println(car2)
}
