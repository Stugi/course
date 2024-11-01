package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const filename = "hi.txt"

/***
Пакет bufio предоставляет структуру Reader для буферизированного чтения данных.
Эта структура может использоваться не только как обычный reader,
она также предоставляет удобные методы для чтения данных, например:

ReadLine()
ReadString(delim byte)
Пример использования ReadLine:
***/

func main() {
	f, _ := os.OpenFile(filename, os.O_RDONLY, 0777)

	fileReader := bufio.NewReader(f)

	for {
		line, _, err := fileReader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}

/**
При вызове данного метода bufio.Reader начинает чтение из источника до первого символа перевода
на новую строку, после этого возвращает строку.

При решении прикладных задач предпочтительно использовать эти методы для уменьшения количества
строк кода.
**/
