package main

import "fmt"

func main() {
	mymap := make(map[int]int)
	mymap[1] = 1
	mymap[2] = 2

	fmt.Println(mymap[3])
	fmt.Println("=====")

	if mymap[3] != 0 {
		fmt.Println(mymap[3])
	}
	fmt.Println("=====")

	// use the second value of the returned value to check if this variable is defined
	//ok?
	if value, ok := mymap[3]; ok {
		fmt.Println(value)
	}
}