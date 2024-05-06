package main

import "fmt"

func main() {
	isReady := true
	memory := 8

	if isReady {
		fmt.Println("Success fetching data!")
	}

	if memory >= 16 {
		fmt.Println("Can use Android Studio :)")
	} else {
		fmt.Println("Oops, your device does not meet the minimum requirements :(")
	}

	name := "Ridwan"

	// short statement
	if length := len(name); length > 6 {
		fmt.Println("Nama terlalu panjang!")
	} else {
		fmt.Println("Nama sudah oke!")
	}
}
