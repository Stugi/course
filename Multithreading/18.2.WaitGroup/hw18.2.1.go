package main

import (
	"fmt"
	"sync"
)

var countGoRoutines int = 5
var printTimes int = 10

func main() {
	// Создайте программу, которая запускает 5 рутин, каждая из которых печатает свой порядковый номер 10 раз.
	// Добиться такого результата за один вызов wg.Add.
	var wg sync.WaitGroup

	wg.Add(countGoRoutines)

	for i := 0; i < countGoRoutines; i++ {
		go func() {
			for j := 0; j < printTimes; j++ {
				fmt.Println(i)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
