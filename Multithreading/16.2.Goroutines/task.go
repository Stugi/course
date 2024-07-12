package main

import (
	"fmt"
	"os"
	"time"
)

//Задание 16.2.1
// Напишите программу, которая запустит n горутин,
// каждая из которых будет один раз в секунду выводить свой идентификатор в консоль.
// Идентификатором считается порядковый номер горутины.
// Число n задаётся пользователем путём ввода числа в консоль.
// Программа должна выполняться до тех пор, пока пользователь не нажмёт клавишу Enter.

func main() {
	var count int

	fmt.Println("Enter gorutines count")
	_, err := fmt.Scan(&count)
	checkerr(err)

	for i := 0; i < count; i++ {
		i := i
		go func() {
			for {
				fmt.Println(i)
				time.Sleep(time.Second)
			}
		}()
	}
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func checkerr(err error) {
	if err != nil {
		fmt.Println("FATAL: ", err.Error())
		os.Exit(1)
	}
}
