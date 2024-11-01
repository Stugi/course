package statistics

import (
	"os"
	"testing"
)

var password string

// Функция TestMain() вызывается каждый раз при запуске go test.
// При наличии TestMain() для выполнения тестов необходимо вызвать в теле TestMain()функцию m.Run(),
// иначе тесты не запустятся.
func TestMain(m *testing.M) {
	password = os.Getenv("password")
	os.Exit(m.Run())
}

func TestAvg(t *testing.T) {
	type args struct {
		nums []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test 1",
			args: args{nums: []float64{1, 2, 3}},
			want: 2.0,
		},
		{
			name: "test 2",
			args: args{nums: []float64{1, 2, 3, 4}},
			want: 2.5,
		},
		{
			name: "test 3",
			args: args{nums: []float64{5, 4, 3, 3, 4, 5}},
			want: 4.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Avg(tt.args.nums); got != tt.want {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxN(t *testing.T) {
	type args struct {
		nums []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test 1",
			args: args{nums: []float64{1, 2, 3}},
			want: 3,
		},
		{
			name: "test 2",
			args: args{nums: []float64{1, 2, 3, 4}},
			want: 4,
		},
		{
			name: "test 3",
			args: args{nums: []float64{5, 4, -3, 3, 4, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxN(tt.args.nums); got != tt.want {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "test 3",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "test 5",
			args: args{n: 5},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Num(tt.args.n); got != tt.want {
				t.Errorf("Num() = %v, want %v", got, tt.want)
			}
		})
	}
}
