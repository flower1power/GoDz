package storage

import (
	"encoding/json"
	"os"

	"3-bin/bins"

	"github.com/fatih/color"
)

const FILE_NAME string = "bin.json"

func SavedBinToFile(bytes []byte) error {
	file, err := os.Create(FILE_NAME)

	if err != nil {
		color.Red("Не удалось создать файл bin.json")
		return err
	}

	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		color.Red("Не удалось записать данные в bin.json")
		return err
	}

	return nil
}

func ReadFileBin() (*bins.Bin, error) {
	file, err := os.ReadFile(FILE_NAME)


	if err != nil {
		color.Red("Не удалось прочитать файл bin.json")
		return nil, err
	}

	var bin bins.Bin

	err = json.Unmarshal(file, &bin)

	if err != nil {
		color.Red("Не удалось преобразовать файл в структуру Bin")
		return nil, err
	}

	return &bin, nil
}

