package main

import (
	"math"
	"net/http"
	_ "net/http/pprof"
)

// Для запуска http://localhost:80/debug/pprof (http://localhost/debug/pprof/)
func main() {
	go compute1(1_000_000, 1_000_000)
	go compute2(2_000_000, 2_000_000)
	go mem1()
	go mem2()
	http.ListenAndServe(":80", nil)
}

func compute1(n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_ = math.Sin(float64(i)) * math.Cos(float64(j))
		}
	}
}
func compute2(n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_ = math.Sin(float64(i)) * math.Cos(float64(j))
		}
	}
}

func mem1() {
	var arr []int
	for i := 0; i < 5_000_000; i++ {
		arr = append(arr, i)
	}
	_ = arr
	select {}
}
func mem2() {
	var arr []int
	for i := 3_000_000; i > 0; i-- {
		arr = append(arr, i)
	}
	_ = arr
	select {}
}
