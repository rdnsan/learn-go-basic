package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string) // type assertions
	fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
