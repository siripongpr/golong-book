package main

import (
	"strconv"
	"fmt"
)

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(number, fizzbuzz(number))
		fmt.Println(number, fizzbuzz2(number))
	}
}

func fizzbuzz(number int) (display string) {
	if number%15 == 0 {
		display = "FizzBuzz"
	} else if number%3 == 0 {
		display = "Fizz"
	} else if number%5 == 0 {
		display = "Buzz"
	} else {
		display = strconv.Itoa(number)
	}
	return
}

func fizzbuzz2(number int) string {
	ln := [3]int{15, 3, 5}
	str := [3]string{"FizzBuzz", "Fizz", "Buzz"}
	for i := 0; i < len(ln); i++ {
		if number%ln[i] == 0 {
			return str[i]
		}
	}
	return strconv.Itoa(number)
}