package main

import "fmt"

func main() {
	x := 123.5
	correct := false
	for i := 0; i < 5 || correct == true; i++ {
		fmt.Print("Enter a number: ")
		var input float64
		fmt.Scanf("%f", &input)
		if input == x {
			correct = true
			fmt.Println("Found")
		} else if input > x {
			fmt.Println("Too high")
		} else {
			fmt.Println("Too low")
		}
	}
	if correct == false {
		fmt.Println("Too much")
	}
}