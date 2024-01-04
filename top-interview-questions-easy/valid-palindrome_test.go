package top_interview_questions_easy

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example 1",
			args{
				s: "A man, a plan, a canal: Panama",
			},
			true,
		},
		{
			"Example 2",
			args{
				s: "race a car",
			},
			false,
		},
		{
			"Example 3",
			args{
				s: "0P",
			},
			false,
		},
		{
			"Example 4",
			args{
				s: " ",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	b.Run("isPalindrome", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isPalindrome("A man, a plan, a canal: Panama")
		}

	})

	b.Run("isPalindrome2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isPalindrome2("A man, a plan, a canal: Panama")
		}
	})

	b.Run("isPalindrome3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isPalindrome3("A man, a plan, a canal: Panama")
		}
	})

	b.Run("isPalindrome4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isPalindrome4("A man, a plan, a canal: Panama")
		}
	})

	b.Run("isPalindrom5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isPalindrome5("A man, a plan, a canal: Panama")
		}
	})
}
