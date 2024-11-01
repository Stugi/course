package statistics

// Avg возвращает среднее значение массива чисел.
func Avg(nums []float64) float64 {
	var sum float64
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}

// MaxN возвращает максимальное значение массива чисел.
func MaxN(nums []float64) float64 {
	var max float64
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

// Num возвращает число Фибоначчи
// по индексу.
func Num(n int) int {
	x1, x2 := 0, 1
	for i := 1; i <= n; i++ {
		x1, x2 = x2, x1+x2
	}
	return x2
}
