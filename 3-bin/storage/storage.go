package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"3-bin/bins"

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
	// file, err := file.ReadFile(FILE_NAME)
	// if err != nil {
	// 	return nil, err
	// }

	ext := filepath.Ext(FILE_NAME)

	if ext != ".json" {
		err := fmt.Errorf("указанный файл не является json: %s", FILE_NAME)
		color.Red(err.Error())
		return nil, err
	}

  isExist, err := exists(FILE_NAME)
	if err != nil {
		color.Red("ошибка при проверке файла: %s", err)
		return nil, err
	}

  if !isExist {
		err := fmt.Errorf("указанный файл не найден: %s", FILE_NAME)
		color.Red(err.Error())
		return nil, err
  }


	file, err := os.ReadFile(FILE_NAME)

	if err != nil {
		color.Red("Не удалось прочитать файл %s", FILE_NAME)
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

func exists(name string) (bool, error) {
	_, err := os.Stat(name)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, fs.ErrNotExist){
		return false, nil
	}

	return false, err
}