package top_interview_questions_easy

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{
				s: "42",
			},
			42,
		},
		{
			"Example 2",
			args{
				s: "   -42",
			},
			-42,
		},
		{
			"Example 3",
			args{
				s: "4193 with words",
			},
			4193,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_myAtoi(b *testing.B) {
	b.Run("myAtoi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			myAtoi("4193 with words")
		}

	})

	b.Run("myAtoi2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			myAtoi2("4193 with words")
		}

	})
}
