package main

import (
	"fmt"
	"os"
	"strings"
)

type ratesMap = map[string]float64

var exchangeRates = ratesMap{
	"USD": 1,
	"EUR": 0.85,
	"RUB": 89.6,
}

const errorCurrencyMsg = "Поддерживается только (USD RUB EUR) - повторите ввод или нажмите 'q' для выхода"
const errorAmountMsg = "Сумма должна быть больше 0 - повторите ввод или нажмите 'q' для выхода"
const fromMsg = "Введите начальную валюту (USD RUB EUR): "
const toMsg = "Введите целевую валюту (USD RUB EUR): "
const amountMsg = "Введите сумму: "

func main() {
	fmt.Println("__Конвертер валют__")

	from := getCurrencyInput(fromMsg, &exchangeRates)
	to := getCurrencyInput(toMsg, &exchangeRates)
	amount := getAmount()

	result := converter(amount, from, to, &exchangeRates)
	fmt.Printf("Итог конвертации: %.2f %s\n", result, to)
}

func getCurrencyInput(prompt string, rates *ratesMap) string {
	for {

		fmt.Print(prompt)
		var currency string
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)

		quit(currency)
		_, ok := (*rates)[currency]

		if ok {
			return currency
		}

		fmt.Println(errorCurrencyMsg)
	}
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

func converter(amount float64, from string, to string, rates *ratesMap) float64 {
	usdAmount := amount / (*rates)[from]
	return usdAmount * (*rates)[to]
}
