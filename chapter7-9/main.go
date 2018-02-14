package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice); i++ {
		fmt.Println(i, slice[i])
	}
	fmt.Println("With Range")
	for i, slice := range slice {
		fmt.Println(i, slice)
	}
}