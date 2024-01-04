package sum_1d_array

import (
	"reflect"
	"testing"
)

func Test_runningSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runningSum2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runningSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_runningSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runningSum([]int{1, 2, 3, 4})
	}
}

func Benchmark_runningSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runningSum2([]int{1, 2, 3, 4})
	}
}

func Benchmark_runningSum3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runningSum3([]int{1, 2, 3, 4})
	}
}
