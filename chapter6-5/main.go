package main

import (
	"strconv"
	"fmt"
)

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(number, fizzbuzz(number))
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
