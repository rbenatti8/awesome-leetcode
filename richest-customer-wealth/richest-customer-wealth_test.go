package richest_customer_wealth

import "testing"

func Test_maximumWealth(t *testing.T) {
	type args struct {
		accounts [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				accounts: [][]int{
					{1, 2, 3},
					{3, 2, 1},
				},
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				accounts: [][]int{
					{1, 5},
					{7, 3},
					{3, 5},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumWealth(tt.args.accounts); got != tt.want {
				t.Errorf("maximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maximumWealth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumWealth([][]int{
			{1, 2, 3},
			{3, 2, 1},
		})
	}
}

func Benchmark_maximumWealth2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumWealth2([][]int{
			{1, 2, 3},
			{3, 2, 1},
		})
	}
}
