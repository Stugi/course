package main

import (
	"fmt"
	"sync"
	"time"
)

//написать код многопоточной программы,
// в которой 8 горутин параллельно осуществляют наращивание целочисленной переменной
//с определённым шагом до какого-то предела

// Перепишите приведённый выше пример со счётчиком из основного текста,
// но вместо примитивов из пакета atomic используйте условную переменную
// и попробуйте реализовать динамическую проверку достижения конечного значения счётчиком.

const step int = 1
const countGorutins int = 8
const interationAmount int = 1000

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	var counter int = 0

	for i := 1; i <= countGorutins; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				c.L.Lock()

				if counter >= interationAmount {
					c.L.Unlock()
					return
				}

				counter += step
				fmt.Println("counter", counter, i)
				c.L.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
