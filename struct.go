package main

import "fmt"

// field must be PascalCase

type Customer struct {
	Name, Address string
	Age           int
}

// struct method

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var frank Customer

	frank.Name = "Frank Martin"
	frank.Address = "San Jose"
	frank.Age = 30

	fmt.Println(frank)
	fmt.Println(frank.Name)

	meong := Customer{
		Name:    "Meong",
		Address: "Cat Island",
		Age:     2,
	}
	fmt.Println(meong)

	monkey := Customer{"Monkey D. Luffy", "East Blue", 18}
	fmt.Println(monkey)

	meong.sayHello("Simba")
}
