package main

import "fmt"

func main() {
	fmt.Println(add(1, 2, 3))

	x5 := []int{1, 2, 3}
	fmt.Println(add(x5...))

	numbers := []int{17, 28, 39}
	for k, v := range numbers {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
