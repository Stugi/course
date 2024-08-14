package main

import (
	"fmt"
	"sync"
)

// Напишите код, в котором несколько горутин увеличивают значение целочисленного счётчика
// и синхронизируют свою работу через канал. Нужно предусмотреть возможность настройки количества
// используемых горутин и конечного значения счётчика, до которого его следует увеличивать.

// Попробуйте реализовать счётчик с элементами ООП (в виде структуры и методов структуры).
// Попробуйте реализовать динамическую проверку достижения счётчиком нужного значения.

const amountOfGoroutins int = 6
const endValue int = 1000

func main() {

	counter := NewCounter()

	for i := 0; i < amountOfGoroutins; i++ {
		fmt.Printf("Горутина %d начала работу\n", i)
		go worker(counter, i, endValue)
	}

	for i := 0; i < amountOfGoroutins; i++ {
		<-counter.quit
	}

	fmt.Println(counter.getValue())
}

type Counter struct {
	value int
	mu    sync.Mutex
	quit  chan bool
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
		quit:  make(chan bool),
	}
}

func (c *Counter) Increase() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) getValue() int {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.value
}

func (c *Counter) Quit() {
	c.quit <- true
}

func worker(c *Counter, id int, target int) {
	for {
		select {
		case <-c.quit:
			return
		default:
			c.Increase()
			if c.getValue() >= target {
				c.Quit()
			}
		}
	}
}
