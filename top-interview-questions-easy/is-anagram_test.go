package top_interview_questions_easy

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example 1",
			args{s: "anagram", t: "nagaram"},
			true,
		},
		{
			"Example 2",
			args{s: "rat", t: "car"},
			false,
		},
		{
			"Example 3",
			args{s: "a", t: "ab"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagram("anagram", "nagaram")
	}
}

func Benchmark_isAnagram2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagram2("anagram", "nagaram")
	}
}
