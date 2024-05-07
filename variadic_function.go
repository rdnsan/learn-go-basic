package main

import "fmt"

// varargs (variable arguments)
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	sum := sumAll(5, 10, 15)
	fmt.Println(sum) // 30

	values := []int{10, 20, 30}
	sumValues := sumAll(values...)

	fmt.Println(sumValues)
}
