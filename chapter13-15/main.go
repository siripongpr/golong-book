package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)

	for i := range ch { // deadlock if channel is still opened because this line of code is waiting for data from channel ch
		fmt.Println(i)
	}
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}