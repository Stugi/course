package main

import (
	"fmt"
	"math/rand"
	"time"
)

//  1/3Задание 17.6.1 (HW-04)
// Напишите код, в котором имеются два канала сообщений из целых чисел, так,
// чтобы приём сообщений из них никогда не приводил к блокировке и
// чтобы вероятность приёма сообщения из первого канала была выше в 2 раза, чем из второго.
// *Если хотите, можете написать код, который бы демонстрировал это соотношение.

func main() {
	rand.Seed(time.Now().Unix())

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		for {
			sendMessage(ch1, 1)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	go func() {
		for {
			sendMessage(ch2, 2)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		}
	}()

	count1 := 0
	count2 := 0
	for i := 0; i < 100; i++ {
		select {
		case message := <-ch1:
			fmt.Println("Получено сообщение из канала 1:", message)
			count1++
		case message := <-ch2:
			fmt.Println("Получено сообщение из канала 2:", message)
			count2++
		default:
			// fmt.Println("Нет сообщений в каналах")
			time.Sleep(10 * time.Millisecond)
		}
	}

	fmt.Println("Количество сообщений в канале 1:", count1)
	fmt.Println("Количество сообщений в канале 2:", count2)

}

func sendMessage(ch chan int, message int) {
	ch <- message
}
