package fizz_buzz

import (
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{n: 3},
			want: []string{"1", "2", "Fizz"},
		},
		{
			name: "Example 2",
			args: args{n: 5},
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "Example 3",
			args: args{n: 15},
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_fizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzz(15)
	}
}

func Benchmark_fizzBuzz2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzz2(15)
	}
}

//func Benchmark_intToString(b *testing.B) {
//	b.Run("intToString", func(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			intToString(15)
//		}
//	})
//
//	b.Run("intToString2", func(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			intToString2(15)
//		}
//	})
//}
