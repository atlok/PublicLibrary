package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateTxtFileName(name string) string {
	return name + ".txt"
}

func GetBookshelfPath() (string, error) {
	executedPath, err := os.Executable()

	if err != nil {
		return "", fmt.Errorf("Ошибка получения директории запуска: %v \n", err)
	}

	libraryPath := filepath.Join(filepath.Dir(executedPath), "library")

	if _, err := os.Stat(libraryPath); os.IsNotExist(err) {
		err = os.Mkdir(libraryPath, 0755)
		if err != nil {
			return "", fmt.Errorf("Ошибка создания директории библиотеки: %v \n", err)
		}
		log.Println("Создана директоория: " + libraryPath)
	}
	return libraryPath, nil
}
