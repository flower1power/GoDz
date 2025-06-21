package main

import "fmt"

const UsdToEur = 0.85
const UsdToRub = 89.6

func main() {
	fmt.Println("__Конвертер валют__")

	firstCurrency, secondCurrency := getUserInput()

	fmt.Println(UsdToRub / UsdToEur)
}

func getUserInput() (string, string) {
	var firstCurrency string
	var secondCurrency string

	fmt.Println("Введите начальную валюту (USD RUB EUR)")
	fmt.Scan(&firstCurrency)

	fmt.Println("Введите в какую валюту конвертировать (USD RUB EUR)")
	fmt.Scan(&secondCurrency)

	return firstCurrency, secondCurrency
}

func converterCurrency(value float64, from string, to string) {
}
