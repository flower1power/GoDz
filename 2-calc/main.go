package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("__Калькулятор__")

	operation, err := setOperations()
	if err != nil {
		return
	}

	numbers, err := setNumbers()
	if err != nil {
		return
	}

	result, err := calculate(operation, numbers)
	if err != nil {
		return
	}

	fmt.Printf("Результат операции: %s равен: %.2f ", operation, result)
}

func setOperations() (string, error) {
	var input string

	for {
		fmt.Print("Введите предполагаемую операцию (AVG - среднее, SUM - сумма, MED - медиана, Q - выход): ")
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Ошибка ввода: ", err)
			continue
		}

		input = strings.ToUpper(strings.TrimSpace(input))

		if input == "Q" {
			return "", fmt.Errorf("выход по запросу пользователя")
		}

		if input == "AVG" || input == "SUM" || input == "MED" {
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

func calculate(operation string, nums []float64) (float64, error) {
	if len(nums) < 1 {
		return 0, nil
	}

	switch operation {
	case "AVG":
		return calculateAVG(nums)
	case "SUM":
		return calculateSum(nums)
	case "MED":
		return calculateMed(nums)
	default:
		return 0, fmt.Errorf("неизвестная операция: %s", operation)
	}
}

func calculateSum(nums []float64) (float64, error) {
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
