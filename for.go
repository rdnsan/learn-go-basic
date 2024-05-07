package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	fmt.Println("Done!")

	for c := 1; c <= 10; c++ {
		fmt.Println("Loop -", c)
	}

	names := []string{"Ridwan", "Nurul", "Ikhsan"}

	for index, name := range names {
		fmt.Println(index, "-", name)
	}

	person := map[string]string{
		"name": "Ridwan",
		"job":  "Full Snack Developer",
	}

	for _, value := range person {
		fmt.Println(value)
	}
}
