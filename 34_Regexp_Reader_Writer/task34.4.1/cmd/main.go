package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	rf "task34.4.1/pkg/reader"
	wf "task34.4.1/pkg/writer"
)

// Напишите программу, которая принимает в качестве аргумента путь директории.
// Программа проверяет список файлов и папок и записывает имена в файл LS.txt.
// Содержимое файла должно быть примерно такое:

//.bash_history (37640 bytes) [FILE]
// .bash_logout (220 bytes) [FILE]
// .bashrc (4608 bytes) [FILE]
// .cache (4096 bytes) [DIRECTORY]
// Desktop (4096 bytes) [DIRECTORY]
// Documents (4096 bytes) [DIRECTORY]
// Downloads (12288 bytes) [DIRECTORY]
// install-nordvpn.sh (101 bytes) [FILE]
// install-sqlite.sh (96 bytes) [FILE]
// some_db.sql (24576 bytes) [FILE]
// some_test.db (8192 bytes) [FILE]

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку данных: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	str = strings.TrimSpace(str)
	fmt.Println("Ваша строка: ", str)
	// можно использовать данные в переменной str
	filesList, err := rf.GetFilesFromPath(str)

	if err != nil {
		fmt.Errorf("GetFilesFromPath() = %v", err)
	}

	err = wf.WriteInfoAboutFiles(filesList)

	if err != nil {
		fmt.Errorf("WriteInfoAboutFiles() = %v", err)
	}

	fmt.Println("Данные записаны в файл")

}
