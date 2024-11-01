package main

import "os"

//Самый просто пример — вывод в консоль.
//Но без использования пакета fmt, к которому вы привыкли.

func main() {
	_, _ = os.Stdout.Write([]byte("Hello, World!!!"))
}
