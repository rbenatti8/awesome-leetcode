package top_interview_questions_easy

import "testing"

func Test_reverseInt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{123},
			321,
		},
		{
			"Example 2",
			args{-123},
			-321,
		},
		{
			"Example 3",
			args{120},
			21,
		},
		{
			"Example 4",
			args{0},
			0,
		},
		{
			"Example 5",
			args{1534236469},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseInt(tt.args.x); got != tt.want {
				t.Errorf("reverseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
