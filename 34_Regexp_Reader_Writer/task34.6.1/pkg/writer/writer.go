package writer

import (
	"bufio"
	"fmt"
	"os"
)

// const fileName = "result.txt"

func WriteInfoAboutFiles(path string, strs []string) error {
	// Открываем файл для записи
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Создаем буферизированный писатель
	writer := bufio.NewWriter(file)

	// Записываем результаты в файл
	for _, str := range strs {

		_, err := writer.WriteString(str + "\n")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// Запускаем флеш-буфер, чтобы записать данные в файл
	writer.Flush()

	return nil
}
