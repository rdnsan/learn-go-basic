package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function!")
}

func run() {
	defer logging()
	fmt.Println("Server running on localhost:3000")
}

func main() {
	run()
}
