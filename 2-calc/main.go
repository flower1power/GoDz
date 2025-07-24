package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type operation = map[string]func([]float64) (float64, error)

var operationMap = operation{
	"SUM": calculateSum,
	"AVG": calculateAVG,
	"MED": calculateMed,
}

func main() {
	fmt.Println("__Калькулятор__")

	operation, err := setOperations(&operationMap)
	if err != nil {
		fmt.Println("Завершение:", err)
		return
	}

	numbers, err := setNumbers()
	if err != nil {
		fmt.Println("Завершение:", err)
		return
	}

	onFunc := operationMap[operation]
	result, err := onFunc(numbers)
	if err != nil {
		fmt.Println("Ошибка при вычислении:", err)
		return
	}

	fmt.Printf("Результат операции: %s равен: %.2f ", operation, result)
}

func setOperations(o *operation) (string, error) {
	var input string

	for {
		fmt.Print("Введите предполагаемую операцию (AVG, SUM, MED, Q - выход): ")
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Ошибка ввода: ", err)
			continue
		}

		input = strings.ToUpper(strings.TrimSpace(input))

		if input == "Q" {
			return "", fmt.Errorf("выход по запросу пользователя")
		}

		_, ok := (*o)[input]
		if ok {
			return input, nil
		}

		fmt.Println("Недопустимая операция. Повторите ввод.")
	}
}

func setNumbers() ([]float64, error) {
	for {
		fmt.Print("Введите числа через (1, 1.2, 3) или 'Q' для выхода: ")
		var input string
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Ошибка чтения строки:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if strings.ToUpper(input) == "Q" {
			return nil, fmt.Errorf("выход по запросу пользователя")
		}

		parts := strings.Split(input, ",")

		var numbers []float64
		valid := true

		for _, part := range parts {
			numStr := strings.TrimSpace(part)
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				fmt.Printf("Ошибка: '%s' не является числом.\n", numStr)
				valid = false
				break
			}
			numbers = append(numbers, num)
		}

		if valid && len(numbers) > 0 {
			return numbers, nil
		}

		fmt.Println("Ввод невалиден. Попробуйте снова.")
	}
}

func calculateSum(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("пустой список чисел")
	}

	var sum float64
	for _, v := range nums {
		sum += v
	}

	return sum, nil
}

func calculateAVG(nums []float64) (float64, error) {
	sum, err := calculateSum(nums)
	if err != nil {
		return 0, err
	}

	return sum / float64(len(nums)), nil
}

func calculateMed(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("пустой список чисел")
	}

	newNums := make([]float64, len(nums))
	copy(newNums, nums)

	sort.Float64s(newNums)

	mid := len(newNums) / 2

	if len(newNums)%2 == 0 {
		return (newNums[mid-1] + newNums[mid]) / 2, nil
	} else {
		return newNums[mid], nil
	}
}
