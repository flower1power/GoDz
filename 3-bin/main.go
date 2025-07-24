package main

import (
	"3-bin/bins"
	"3-bin/storage"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("__BIN__")

Menu:
	for {
		variant := promptData([]string{
			"1. Создать",
			"2. Выход",
			"Выберите вариант",
		})

		switch variant {
		case "1":
			_ = storage.Saved(createBin())
		case "2":
			break Menu
		default:
			fmt.Println("Выберите любой из 2 вариантов")
		}
	}

}

func createBin() *bins.Bin {
	var isPrivate bool
	var name string

	fmt.Print("Приватный bin(0/1)?: ")
	_, err := fmt.Scanln(&isPrivate)

	if err != nil {
		color.Red("Нe удалось преобразовать в boolean")
	}

	fmt.Print("Имя bin: ")
	_, err = fmt.Scanln(&isPrivate)

	if err != nil {
		color.Red("Нe удалось преобразовать в строку")
	}

	return bins.NewBin(isPrivate, name)
}

func promptData[T any](prompts []T) string {
	var input string

	for i, line := range prompts {
		if i == len(prompts)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}

	_, err := fmt.Scanln(&input)
	if err != nil {
		color.Red("Нe удалось преобразовать в строку")
	}

	return input
}
