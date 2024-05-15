package main

import "fmt"

type TechStack struct {
	FrontEnd, BackEnd string
}

func main() {
	techStack1 := TechStack{FrontEnd: "TypeScript", BackEnd: "Ruby"}
	techStack2 := &techStack1 // pointer (reference)

	techStack2.BackEnd = "Go"

	fmt.Println(techStack1)
	fmt.Println(techStack2)

	// techStack2 = &TechStack{"JavaScript", "Java"}
	*techStack2 = TechStack{"Dart", "Laravel"} // asterisk operator (*)

	fmt.Println(techStack1)
	fmt.Println(techStack2)
}
