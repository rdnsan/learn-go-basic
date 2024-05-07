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

	name := "Ridwan"
	// short statement
	switch length := len(name); length > 6 {
	case true:
		fmt.Println("Nama terlalu panjang!")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
