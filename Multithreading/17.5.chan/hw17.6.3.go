package main

import (
	"fmt"
	"sync"
)

//  3/3Задание 17.6.3 (HW-04)
// Напишите программу, которая делает следующее: одна горутина по порядку отсылает числа от 1 до 100 в канал,
// вторая горутина их принимает в правильном порядке и печатает на экран (в консоль).

func main() {

	wg := sync.WaitGroup{}
	chan1 := make(chan int)

	f1 := func() {
		fmt.Println("Первая горутина")
		defer wg.Done()
		defer close(chan1)
		for i := 1; i <= 100; i++ {
			chan1 <- i
		}
	}

	f2 := func() {
		fmt.Println("Вторая горутина")
		defer wg.Done()
		for c := range chan1 {
			fmt.Println(c)
		}
	}
	wg.Add(2)
	go f1()
	go f2()
	wg.Wait()
}
