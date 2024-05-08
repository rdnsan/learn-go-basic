package main

import "fmt"

// empty interface {} | any

func Ups() interface{} {
	return 1
}

func main() {
	var empty any = Ups()
	fmt.Println(empty)
}
