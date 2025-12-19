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
	var error error

	for {
		sourceCurrency, error = getTheSourceCurrency()
		if error != nil {
			fmt.Printf("%s\n", error)
			continue
		} else {
			break
		}
	}
	for {
		destCurrency, error = getTheDestinationCurrency(sourceCurrency)
		if error != nil {
			fmt.Printf("%s\n", error)
			continue
		} else {
			break
		}
	}
	for {
		amount, error = getTheAmount(sourceCurrency)
		if error != nil {
			fmt.Printf("%s\n", error)
			continue
		} else {
			break
		}
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

func getTheSourceCurrency() (string, error) {
	var sourceCurrency string
	fmt.Print("Введите исходную валюту (USD, RUB, EURO): ")
	fmt.Scan(&sourceCurrency)
	if sourceCurrency != "USD" && sourceCurrency != "RUB" && sourceCurrency != "EURO" {
		return "", errors.New("Используйте только предложенные варианты валют")
	}
	return sourceCurrency, nil
}

func getTheDestinationCurrency(currency string) (string, error) {
	var destCurrency string
	fmt.Print("Введите желаемую валюту (USD, RUB, EURO): ")
	fmt.Scan(&destCurrency)
	if destCurrency == currency {
		return "", errors.New("Валюта должна различаться")
	}
	if destCurrency != "USD" && destCurrency != "RUB" && destCurrency != "EURO" {
		return "", errors.New("Используйте только предложенные варианты валют")
	}
	return destCurrency, nil
}

func getTheAmount(sourceCurrency string) (float64, error) {
	var amount float64
	fmt.Printf("Введите количество %s: ", sourceCurrency)
	fmt.Scan(&amount)
	if amount < 0 {
		return 0, errors.New("Значение не может быть меньше нуля")
	}
	return amount, nil
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
		return value * USD_TO_EUR, nil
	}
	if sourceCurrency == "RUB" && destCurrency == "EURO" {
		return value / EURO_TO_RUB, nil
	}
	if sourceCurrency == "RUB" && destCurrency == "USD" {
		return value / USD_TO_RUB, nil
	}
	return 0, errors.New("Неизвестная ошибка")
}
