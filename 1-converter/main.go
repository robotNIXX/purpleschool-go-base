package main

import "fmt"

func main() {
	const USD_TO_EUR = 0.86
	const USD_TO_RUB = 77.48
	const EURO_TO_RUB = USD_TO_RUB / USD_TO_EUR

	var rub float64
	amount := getTheCurrencyValues()
	rub = (amount * EURO_TO_RUB)
	fmt.Printf("%.2f EURO is %.2f RUB\n", amount, rub)
}

/*
Get the euro amount
*/
func getTheCurrencyValues() (amount float64) {
	fmt.Print("Введите количество евро: ")
	fmt.Scan(&amount)

	return
}

func convertTheCurrency(value float64, sourceCurrency string, destCurrency string) {
	// some logic
}
