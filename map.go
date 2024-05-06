package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Ridwan",
		"job":  "Developer",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["job"])

	book := make(map[string]string)

	book["title"] = "Clean Code"
	book["author"] = "Robert C Martin"
	book["year"] = "2008"

	fmt.Println(book)

	delete(book, "year")

	fmt.Println(book)
}
