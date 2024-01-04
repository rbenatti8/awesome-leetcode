package top_interview_questions_easy

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{s: "leetcode"},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{s: "loveleetcode"},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{s: "aabb"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar3(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_firstUniqChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstUniqChar("loveleetcode")
	}
}

func Benchmark_firstUniqChar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstUniqChar2("loveleetcode")
	}
}

func Benchmark_firstUniqChar3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstUniqChar3("loveleetcode")
	}
}
