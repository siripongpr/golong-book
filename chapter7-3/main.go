package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr)

	slice := arr[0:3]
	fmt.Println(slice)
}