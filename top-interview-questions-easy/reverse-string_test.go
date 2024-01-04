package top_interview_questions_easy

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "example 1",
			args: args{
				s: []byte{'h', 'e', 'l', 'l', 'o'},
			},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_reverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	}
}

func Benchmark_reverseString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseString2([]byte{'h', 'e', 'l', 'l', 'o'})
	}
}
