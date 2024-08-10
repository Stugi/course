package main

import (
	"fmt"
	"time"
)

// 2/3Задание 17.6.2 (HW-04)
// Напишите код, в котором имеются два канала сообщений из целых чисел так,
// чтобы приём сообщений всегда приводил к блокировке.
// Приёмом сообщений из обоих каналов будет заниматься главная горутина.
//  Сделайте так, чтобы во время такого «бесконечного ожидания» сообщений выполнялась
//фоновая работа в виде вывода текущего времени в консоль.

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	var ticker *time.Ticker = time.NewTicker(time.Second * 1)
	var t time.Time
	for {
		select {
		case <-chan1:
			fmt.Println("Первый канал")
		case <-chan2:
			fmt.Println("Второй канал")
		default:
			t = <-ticker.C
			outputMessage := []byte("Время: ")
			// Метод AppendFormat преобразует объект time.Time
			// к заданному строковому формату (второй аргумент)
			// и добавляет полученную строку к строке, переданной в первом
			// аргументе
			outputMessage = t.AppendFormat(outputMessage, "15:04:05")
			fmt.Println(string(outputMessage))
		}

	}
}
