package storage

import (
	"encoding/json"
	"os"

	"3-bin/bins"
	"3-bin/file"

	"github.com/fatih/color"
)

const FILE_NAME string = "bin.json"

func SavedBinToFile(bin bins.Bin) error {
	binList, err := ReadFileBin()
	if err != nil {
		return err
	}

	bins.AddBinList(binList, &bin)

	file, err := os.Create(FILE_NAME)

	if err != nil {
		color.Red("Не удалось создать файл bin.json")
		return err
	}

	defer file.Close()

	b, err := json.Marshal(binList)
	if err != nil {
		color.Red("Не удалось преобразовать bin в []byte")
		return err
	}

	_, err = file.Write(b)
	if err != nil {
		color.Red("Не удалось записать данные в bin.json")
		return err
	}

	return nil
}

func ReadFileBin() (*bins.BinList, error) {
	file, err := file.ReadFile(FILE_NAME)
	if err != nil {
		return nil, err
	}

	var binList bins.BinList

	err = json.Unmarshal(file, &binList)

	if err != nil {
		color.Red("Не удалось преобразовать файл в структуру Bin")
		return nil, err
	}

	return &binList, nil
}

