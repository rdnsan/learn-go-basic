package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Ridwan"
	names[1] = "Nurul"
	names[2] = "Ikhsan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	otherNames := [2]string {
		"Frank",
		"Martin",
	}

	fmt.Println(otherNames)

	nums := [3]int {
		77,
		96,
	}

	fmt.Println(nums) // index 2 = 0 (default value for int)

	level := [...]int { 1, 2, 3, 4, 5 }

	fmt.Println(level)
	fmt.Println(len(level))

	level[2] = 666 // change | update
	fmt.Println(level)
}
