package main

import (
	"fmt"
	"time"
)

func main() {
	volumn := 10
	start := time.Now()

	//container := order(volumn)
	container := order2(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

/*func order(volumn int) (container []string) {
	for i := 1; i <= volumn; i++ {
		// cashier receive order
		time.Sleep(5 * time.Millisecond)
		coffee := fmt.Sprintf("order: %d", i)

		// barista brew coffee
		time.Sleep(100 * time.Millisecond)
		coffee = fmt.Sprintf("%s %s", coffee, "espresso")

		// waiter serve cofee
		time.Sleep(5 * time.Millisecond)
		container = append(container, fmt.Sprintf("%s %s", coffee, "ready :)"))
	}

	return
}*/

func order2(volumn int) (container []string) {
	order := make(chan string, 2)
	brew := make(chan string, 40)
	serve := make(chan string, 2)

	for i := 1; i <= volumn; i++ {
		go receiveorder(order, i)
	}

	go brewcoffee(order, brew)
	go servecoffee(brew, serve)

	for coffee := range serve {
		container = append(container, coffee)
	}

	return
}

func receiveorder(out chan<- string, number int) {
	// cashier receive order
	time.Sleep(5 * time.Millisecond)
	coffee := fmt.Sprintf("order: %d", number)
	out <- coffee
}

func brewcoffee(in <-chan string, out chan<- string) {
	for coffee := range in {
		// barista brew coffee
		time.Sleep(100 * time.Millisecond)
		coffee = fmt.Sprintf("%s %s", coffee, "espresso")
		out <- coffee
	}
	close(out)
}

func servecoffee(in <-chan string, out chan<- string) {
	for coffee := range in {
		// waiter serve cofee
		time.Sleep(5 * time.Millisecond)
		coffee = fmt.Sprintf("%s %s", coffee, "ready :)")
		out <- coffee
	}
	close(out)
}