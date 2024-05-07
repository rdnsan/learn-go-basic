package main

import "fmt"

func goodBye(name string) string {
	return "Good bye, " + name
}

func main() {
	sayGoodBye := goodBye
	fmt.Println(sayGoodBye("Frank"))
}
