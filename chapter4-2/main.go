package main

import "fmt"

func main() {
	fmt.Print("Enter a number in Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	//fmt.Println("====")
	output := (input - 32) * 5 / 9
	//fmt.Println(output)
	fmt.Printf("Number in Celsius is %.2f\n", output)
}