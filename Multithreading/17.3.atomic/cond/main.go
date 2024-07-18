package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Перепишите приведённый выше пример со счётчиком из основного текста,
// но вместо примитивов из пакета atomic используйте условную переменную
// и попробуйте реализовать динамическую проверку достижения конечного значения счётчиком.

const step int64 = 1
const interationAmount int = 1000

func main() {
	var counter int64 = 0
	var c = sync.NewCond(&sync.Mutex{})
	increment := func(i int) {
		atomic.AddInt64(&counter, step)
		if i == interationAmount {
			c.Signal()
		}
	}
	for i := 1; i <= interationAmount; i++ {
		go increment(i)
	}
	c.L.Lock()
	c.Wait()
	c.L.Unlock()
	fmt.Println(counter)
}
