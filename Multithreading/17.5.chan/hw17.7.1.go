package main

import (
	"fmt"
	"sync"
)

const amountOfGoroutins int = 6
const endValue int = 2333

func main() {

	counter := NewCounter()

	for i := 0; i < amountOfGoroutins; i++ {
		fmt.Printf("Горутина %d начала работу\n", i)
		go worker(counter, i, endValue)
	}

	<-counter.done
	fmt.Println(counter.getValue())
}

type Counter struct {
	value    int
	mu       sync.Mutex
	quit     chan bool
	done     chan bool
	doneFlag bool
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
		quit:  make(chan bool),
		done:  make(chan bool),
	}
}

func (c *Counter) Increase() {
	defer c.mu.Unlock()
	c.mu.Lock()
	if c.value >= endValue {
		c.doneFlag = true
		return
	}
	c.value++
}

func (c *Counter) getValue() int {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.value
}

func (c *Counter) Quit() {
	c.quit <- true
}

func (c *Counter) isDone() bool {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.doneFlag
}

func worker(c *Counter, id int, target int) {
	for {
		select {
		case <-c.quit:
			return
		default:
			c.Increase()
			if c.isDone() {
				c.done <- true
				return
			}
		}
	}
}
