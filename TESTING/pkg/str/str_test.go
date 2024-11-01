package str

import "testing"

func TestConcat(t *testing.T) {
	s1 := "hello"
	s2 := "world"
	want := "helloworld"
	got := Concat(s1, s2)
	if got != want {
		t.Errorf("Concat() = %v, want %v", got, want)
	}
}

func TestInvert(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Пример 1",
			args: args{
				s: "123456",
			},
			want: "654321",
		},
		{
			name: "Пример 2",
			args: args{
				s: "12345",
			},
			want: "54321",
		},
		{
			name: "Пример 3",
			args: args{
				s: "",
			},
			want: "",
		},

		{
			name: "Пример 4",
			args: args{
				s: "123Пример",
			},
			want: "ремирП321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Invert(tt.args.s); got != tt.want {
				t.Errorf("Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}
