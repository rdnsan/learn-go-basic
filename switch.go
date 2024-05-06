package main

import "fmt"

func main() {
	os := "linux"

	switch os {
	case "linux":
		fmt.Println("Linux User")
	default:
		fmt.Println("Windows User")
	}
}
