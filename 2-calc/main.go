package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var operation string
	var numbers []int
	var error error
	for {
		operation, error = selectOP()
		if error != nil {
			fmt.Println(error)
		} else {
			break
		}
	}
	for {
		numbers, error = yourNumbers()
		if error != nil {
			fmt.Println(error)
		} else {
			break
		}
	}
	switch operation {
	case "AVG":
		fmt.Printf("Average: %d\n", average(numbers))
	case "SUM":
		fmt.Printf("Sum: %d\n", sum(numbers))
	case "MED":
		fmt.Printf("Median: %d\n", median(numbers))
	}
}

func average(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum / len(numbers)
}

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func median(numbers []int) int {
	sort.Ints(numbers)
	return numbers[len(numbers)/2]
}

func selectOP() (string, error) {
	var operation string
	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	fmt.Scan(&operation)
	if operation != "AVG" && operation != "SUM" && operation != "MED" {
		return "", errors.New("Неизвестная операция")
	}
	return operation, nil
}

func yourNumbers() ([]int, error) {
	var enterString string
	var numberInString []string
	fmt.Print("Введите числа через запятую (числа только целые): ")
	fmt.Scan(&enterString)
	numberInString = strings.Split(enterString, ",")
	numbers := make([]int, len(numberInString))
	for i := range numberInString {
		el, errorItem := strconv.Atoi(numberInString[i])
		if errorItem != nil {
			return []int{}, errors.New("Необходимо вводить числа")
		}
		numbers[i] = el
	}

	return numbers, nil
}
