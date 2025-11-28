package main

import "fmt"

func main() {
	const USD_TO_EUR = 0.86
	const USD_TO_RUB = 77.48
	const EURO_TO_RUB = USD_TO_RUB / USD_TO_EUR

	var amount float64
	var rub float64
	fmt.Print("Введите количство евро: ")
	fmt.Scan(&amount)
	rub = (amount * EURO_TO_RUB)
	fmt.Printf("%.2f EURO is %.2f RUB\n", amount, rub)
}
