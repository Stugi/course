package main

//Требуется доработать текущую версию приложения, чтобы echo сервер мог принимать произвольные данные, не только целые числа.

import (
	"fmt"
	"log"
)

func main() {
	var input string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели: %s\n", input)
}
