package main

import "fmt"

type NewVendingMachine struct {
	t int
	f int
	tw int
	o int
	softDrink int
	cannedCoffee int
	drinkingWater int
}

func (vm NewVendingMachine) InsertCoin(coin string) {
	if coin == "T" {
		vm.t++
	} else if coin == "F" {
		vm.f++
	} else if coin == "TW" {
		vm.tw++
	} else if coin == "O" {
		vm.o++
	}
}

func (vm NewVendingMachine) GetInsertedMoney() int {
	return ((vm.t * 10) + (vm.f * 5) + (vm.tw * 2) + vm.o)
}

func main() {
	vm := NewVendingMachine(Coins, Items)

	// Buy SD(soft drink) with exact change
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney())	// 18
	can := vm.SelectSD()
	fmt.Println(can) // SD

	// Buy CC(canned coffee) without exact change
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney())	// 20
	can = vm.SelectCC()
	fmt.Println(can) // CC, F, TW, O

	// Start adding change but hit coin return
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney())	// 25
	change := vm.CoinReturn()
	fmt.Println(change) // T, T, F
}