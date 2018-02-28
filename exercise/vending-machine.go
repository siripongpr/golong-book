package main

import (
	"fmt"
	"github.com/siripongpr/vendingMachine"
)

var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
var items = map[string]int{"SD": 18, "CC": 12, "DW": 7}

func main() {
	vm := vendingMachine.NewVendingMachine(coins, items)
	fmt.Println("Inserted Money:", vm.InsertedMoney())

	// Buy SD(soft drink) with exact change
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.InsertedMoney())	// 18
	can := vm.SelectSD()
	fmt.Println(can) // SD

	// Buy CC(canned coffee) without exact change
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())	// 20
	can = vm.SelectCC()
	fmt.Println(can) // CC, F, TW, O

	// Start adding change but hit coin return
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.InsertedMoney())	// 25
	change := vm.CoinReturn()
	fmt.Println(change) // T, T, F
}