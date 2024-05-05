package main

import "fmt"

func main() {
	type NoKTP string

	const ktpRidwan NoKTP = "1234567"

	var str string = "7654321"
	example := NoKTP(str)

	fmt.Println(ktpRidwan)
	fmt.Println(example)
}
