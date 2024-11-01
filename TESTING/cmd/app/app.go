package main

import (
	"app/testing/pkg/statistics"
	"fmt"
)

func main() {
	// верный результат
	nums := []float64{1, 2, 3}
	want := 2.0
	got := statistics.Avg(nums)
	if got != want {
		fmt.Printf("получено %f, ожидалось %f\n", got, want)
	}
	fmt.Printf("результат: %f\n", got)
	// ошибка
	nums = []float64{1, 2, 3}
	want = 3.0
	got = statistics.Avg(nums)
	if got != want {
		fmt.Printf("получено %f, ожидалось %f\n", got, want)
	}
	fmt.Printf("результат: %f\n", got)
}
