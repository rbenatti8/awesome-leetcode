package ransom_note

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
		},
		{
			name: "example 2",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
		{
			name: "example 4",
			args: args{
				ransomNote: "aahgs",
				magazine:   "aabblkhisgt",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct3(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_canConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canConstruct("aahgs", "aabblkhisyt")
	}
}

func Benchmark_canConstruct2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canConstruct2("aahgs", "aabblkhisyt")
	}
}

func Benchmark_canConstruct3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canConstruct3("aahgs", "aabblkhisyt")
	}
}

func Benchmark_canConstruct4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canConstruct4("aahgs", "aabblkhisyt")
	}
}
