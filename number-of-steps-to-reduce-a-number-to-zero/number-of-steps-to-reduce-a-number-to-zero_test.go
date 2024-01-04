package number_of_steps_to_reduce_a_number_to_zero

import "testing"

func Test_numberOfSteps(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "Example 1",
		//	args: args{num: 14},
		//	want: 6,
		//},
		//{
		//	name: "Example 2",
		//	args: args{num: 8},
		//	want: 4,
		//},
		{
			name: "Example 3",
			args: args{num: 123},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSteps4(tt.args.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_numberOfSteps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberOfSteps(123)
	}
}

func Benchmark_numberOfSteps2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberOfSteps2(123)
	}
}

func Benchmark_numberOfSteps3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberOfSteps3(123)
	}
}

func Benchmark_numberOfSteps4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberOfSteps4(123)
	}
}
