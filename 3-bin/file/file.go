package file

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func ReadFile(name string) ([]byte, error) {
	ext := filepath.Ext(name)

	if ext != ".json" {
		err := fmt.Errorf("указанный файл не является json: %s", name)
		color.Red(err.Error())
		return nil, err
	}

  isExist, err := existsFile(name)
	if err != nil {
		color.Red("ошибка при проверке файла: %s", err)
		return nil, err
	}

  if !isExist {
		err := fs.ErrNotExist
		color.Red("указанный файл не найден: %s", name)
		return nil, err
  }

	file, err := os.ReadFile(name)

	if err != nil {
		color.Red("Не удалось прочитать файл %s", name)
		return nil, err
	}

	return file, nil
}

func existsFile(name string) (bool, error) {
	_, err := os.Stat(name)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, fs.ErrNotExist){
		return false, nil
	}

	return false, err
}