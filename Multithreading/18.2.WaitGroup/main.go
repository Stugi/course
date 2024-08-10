package main

import (
	"fmt"
	"sync"
	"time"
)

// fib — это функция, расчитывающая N-ное число последовательноси Фибоначчи
func fib(number uint) uint {
	if number == 0 || number == 1 {
		return number
	}

	return fib(number-2) + fib(number-1)
}

func main() {
	computations := []uint{
		34,
		5,
		12,
		25,
		30,
		42,
		3,
	}

	var wg sync.WaitGroup

	n := time.Now()
	results := make([]uint, len(computations))
	for i := range computations {
		wg.Add(1) // всегда снаружи рутины
		go func() {
			results[i] = fib(computations[i])
			wg.Done() // всегда внутри рутины
		}()
	}

	//time.Sleep(2 * time.Second)

	wg.Wait()
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Println("processed in", time.Since(n))
}
