package main

import "fmt"

func main() {
	// Long Declaration
	var x string = "Hello, world"
	fmt.Println(x)

	var y string
	y = "Hello, world"
	fmt.Println(y)

	// Short Declaration
	// Type Inference
	z := "Hello, world"
	//z := 349388493938
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)

	const xx string = "Hello, world"
	//xx = "Other string"

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c)
	fmt.Println("=====")

	v1, v2 := "first", "second"
	v1, v2 = v2, v1
	fmt.Println(v1, v2)
	fmt.Printf("%v\n%v", v1, v2)
}
