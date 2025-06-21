package main

import (
	"fmt"
	"os"
	"strings"
)

const UsdToEur = 0.85
const UsdToRub = 89.6
const errorCurrencyMsg = "Поддерживается только (USD RUB EUR) - повторите ввод или нажмите 'q' для выхода"
const errorAmountMsg = "Сумма должна быть больше 0 - повторите ввод или нажмите 'q' для выхода"
const fromMsg = "Введите начальную валюту (USD RUB EUR): "
const toMsg = "Введите целевую валюту (USD RUB EUR): "
const amountMsg = "Введите сумму: "

func main() {
	fmt.Println("__Конвертер валют__")

	from := getFromCurrency()
	to := getToCurrency()
	amount := getAmount()

	result := convertedCurrency(amount, from, to)
	fmt.Printf("Итог конвертации: %.2f %s\n", result, to)
}

func getFromCurrency() string {
	fmt.Print(fromMsg)
	return getCurrency()
}

func getToCurrency() string {
	fmt.Print(toMsg)
	return getCurrency()
}

func getAmount() float64 {
	fmt.Print(amountMsg)

	for {
		var input string
		fmt.Scan(&input)

		quit(input)

		var amount float64

		_, err := fmt.Sscan(input, &amount)

		if err == nil && amount > 0 {
			return amount
		}

		fmt.Println(errorAmountMsg)
	}
}

func quit(str string) {
	if strings.ToLower(str) == "q" {
		fmt.Println("Выход из программы")
		os.Exit(0)
	}
}

func getCurrency() string {
	for {
		var currency string
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)

		quit(currency)

		if isCurrency(currency) {
			return currency
		}

		fmt.Println(errorCurrencyMsg)
	}
}

func isCurrency(currency string) bool {
	return currency == "EUR" || currency == "USD" || currency == "RUB"
}

func convertedCurrency(amount float64, from string, to string) float64 {
	usd := convertedToUsd(amount, from)

	switch to {
	case "EUR":
		return usd * UsdToEur
	case "RUB":
		return usd * UsdToRub
	}

	return usd
}

func convertedToUsd(amount float64, from string) float64 {
	if from == "USD" {
		return amount
	}

	var usd float64

	switch from {
	case "EUR":
		usd = amount / UsdToEur
	case "RUB":
		usd = amount / UsdToRub
	}

	return usd
}
