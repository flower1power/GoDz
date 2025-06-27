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

	fmt.Print("Введите предполагаемую операцию (AVG - среднее, SUM - сумма, MED - медиана): ")
	_, err := fmt.Scan(&input)

	if err != nil {
		return "", err
	}

	return strings.ToUpper(input), nil
}

func setNumbers() ([]float64, error) {
	var input string
	var numbers []float64

	fmt.Print("Введите числа через (,): ")
	_, err := fmt.Scan(&input)

	if err != nil {
		return numbers, err
	}

	input = strings.TrimSpace(input)

	for _, v := range strings.Split(input, ",") {
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return numbers, err
		}

		numbers = append(numbers, i)
	}

	return numbers, nil
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
		return 0, fmt.Errorf("unknown operation: %s", operation)
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
	newNums := nums

	sort.Float64s(newNums)

	mid := len(newNums) / 2

	if len(newNums)%2 == 0 {
		return (newNums[mid-1] + newNums[mid]) / 2, nil
	} else {
		return newNums[mid], nil
	}
}
