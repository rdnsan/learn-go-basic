package main

import "fmt"

type Address struct {
	City, Province, Country string
}

type TechStack struct {
	FrontEnd, BackEnd string
}

func main() {
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := address1 // copy value (not reference)

	address2.City = "Sumedang"
	fmt.Println(address1)
	fmt.Println(address2)

	techStack1 := TechStack{FrontEnd: "TypeScript", BackEnd: "Ruby"}
	techStack2 := &techStack1 // pointer (reference)

	techStack2.BackEnd = "Go"

	fmt.Println(techStack1)
	fmt.Println(techStack2)
}
