package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.Open("./hi.txt")
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 10)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}

		fmt.Println("> ", string(buf[:n]))
	}
}

// Данная программа выводит список файлов из рабочей директории.
// Показывает имя файла и значение типа bool: директория это (true) или файл (false).
// Если вам потребуется рекурсивная функция для обхода всех сложенных файлов и папок,
// советую использовать функцию filepath.Walk.
func readDir() {

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}

}

func readDirRecursive() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
