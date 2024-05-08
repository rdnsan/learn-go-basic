package main

import "fmt"

type FilterCallback func(string) string

func sayHelloWithFilter(name string, filter FilterCallback) {
	filteredName := filter(name)
	fmt.Println("Hello,", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}
	return name
}

func main() {
	sayHelloWithFilter("Meong", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
