package main

import (
	"errors"
	"fmt"
)

const USD_TO_EUR = 0.86
const USD_TO_RUB = 77.48
const EURO_TO_RUB = USD_TO_RUB / USD_TO_EUR

func main() {

	var sourceCurrency string
	var destCurrency string
	var result float64
	var amount float64

	sourceCurrency, amount, destCurrency, error := getTheCurrencyValues()
	if error != nil {
		fmt.Printf("%s\n", error)
		return
	}
	result, error = convertTheCurrency(amount, sourceCurrency, destCurrency)
	if error != nil {
		fmt.Printf("%s\n", error)
		return
	}
	fmt.Printf("%.2f %s is %.2f %s\n", amount, sourceCurrency, result, destCurrency)
}

/*
Get the euro amount
*/
func getTheCurrencyValues() (string, float64, string, error) {
	var sourceCurrency string
	var destCurrency string
	var amount float64
	fmt.Print("Введите исходную валюту (USD, RUB, EURO): ")
	fmt.Scan(&sourceCurrency)
	if sourceCurrency != "USD" && sourceCurrency != "RUB" && sourceCurrency != "EURO" {
		return "", 0, "", errors.New("Используйте только предложенные варианты валют")
	}
	fmt.Print("Введите желаемую валюту (USD, RUB, EURO): ")
	fmt.Scan(&destCurrency)
	if destCurrency != sourceCurrency {
		if destCurrency != "USD" && destCurrency != "RUB" && destCurrency != "EURO" {
			return "", 0, "", errors.New("Используйте только предложенные варианты валют")
		}
	} else {
		return "", 0, "", errors.New("Используйте только предложенные варианты валют (исключая уже введенную)")
	}
	fmt.Printf("Введите количество %s: ", sourceCurrency)
	fmt.Scan(&amount)
	if amount < 0 {
		return "", 0, "", errors.New("Значение не может быть меньше нуля")
	}

	return sourceCurrency, amount, destCurrency, nil
}

func convertTheCurrency(value float64, sourceCurrency string, destCurrency string) (float64, error) {
	if sourceCurrency == destCurrency {
		return 0, errors.New("Валюта должна различаться")
	}

	if sourceCurrency == "USD" && destCurrency == "RUB" {
		return value * USD_TO_RUB, nil
	}
	if sourceCurrency == "EURO" && destCurrency == "RUB" {
		return value * EURO_TO_RUB, nil
	}
	if sourceCurrency == "USD" && destCurrency == "EURO" {
		return value / USD_TO_EUR, nil
	}
	if sourceCurrency == "RUB" && destCurrency == "EURO" {
		return value / USD_TO_RUB, nil
	}
	if sourceCurrency == "RUB" && destCurrency == "USD" {
		return value / USD_TO_RUB, nil
	}
	return 0, errors.New("Неизвестная ошибка")
}
