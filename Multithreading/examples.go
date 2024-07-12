// Задание 16.2.1
// Напишите программу, которая запустит n горутин,
// каждая из которых будет один раз в секунду выводить свой идентификатор в консоль.
// Идентификатором считается порядковый номер горутины.
// Число n задаётся пользователем путём ввода числа в консоль.
// Программа должна выполняться до тех пор, пока пользователь не нажмёт клавишу Enter.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var count int

	fmt.Println("Enter gorutines count")
	_, err := fmt.Scan(&count)
	checkerr(err) // Огромный минус, если студенты не обрабатывают ошибки. Ожидается что у студента везде будет if err != nil

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

//Задание 16.5.1
// Улучшите InMemoryCache из модуля 14, используя sync.Mutex, сделав его потокобезопасным.
package main

import (
	"sync"
	"time"
)

var _ Cache = &InMemoryCache{}

type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type InMemoryCache struct {
	expireIn time.Duration

	mu      sync.Mutex
	storage map[string]*CacheEntry
}

func (i *InMemoryCache) Set(key string, value interface{}) {
	entry := &CacheEntry{
		settledAt: time.Now(),
		value:     value,
	}
	i.mu.Lock()
	i.storage[key] = entry
	i.mu.Unlock()
}

func (i *InMemoryCache) Get(key string) interface{} {
	i.mu.Lock()
	entry := i.storage[key]
	i.mu.Unlock()
	if entry == nil || time.Since(entry.settledAt) > i.expireIn {
		return nil
	}
	return entry.value
}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		expireIn: expireIn,
		storage: make(map[string]CacheEntry);
	}
}

//Задание 16.5.2
// Напишите программу, которая будет каждую секунду увеличивать значение переменной на единицу и 
//каждую ⅕ секунды выводить в консоль значение этой переменной. 
//Программа должна выполняться n секунд и после этого закрываться.
// n вводит пользователь после запуска программы.

package main

import (
    "fmt"
    "os"
    "time"
)


func main() {
    var n int
    fmt.Print("Ведите длительность работы программы в секундах: ")
    fmt.Fscan(os.Stdin, &n)

    counter := 0

    go func() {
        for {
            time.Sleep(time.Second)
            counter++
        }
    }()

    go func() {
        for {
            fmt.Printf("%v - %v\n", time.Now().Format("15:04:05.00"), counter)
            time.Sleep(time.Millisecond * time.Duration(200))
        }
    }()

    time.Sleep(time.Second * time.Duration(n))

	//Задание 16.6.1
// Улучшите InMemoryCache из задания с прошлого юнита, используя sync.RWMutex.
package main

import (
	"sync"
	"time"
)

var _ Cache = &InMemoryCache{}

type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type InMemoryCache struct {
	expireIn time.Duration

	mu      sync.RWMutex
	storage map[string]*CacheEntry
}

func (i *InMemoryCache) Set(key string, value interface{}) {
	entry := &CacheEntry{
		settledAt: time.Now(),
		value:     value,
	}
	i.mu.Lock()
	i.storage[key] = entry
	i.mu.Unlock()
}

func (i *InMemoryCache) Get(key string) interface{} {
	i.mu.RLock()
	entry := i.storage[key]
	i.mu.RUnlock()
	if entry == nil || time.Since(entry.settledAt) > i.expireIn {
		return nil
	}
	return entry.value
}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		expireIn: expireIn, 
		storage: make(map[string]CacheEntry);
    }
}