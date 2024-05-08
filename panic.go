package main

import "fmt"

func endApp() {
	fmt.Println("Process exit with code 0")
	err := recover()
	fmt.Println(err)
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("Ups something went wrong :(")
	}
}

func main() {
	runApp(true)
	fmt.Println("Dont panic :D")
}
