package sum_1d_array

func sum(nums []int, previousSum int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	s := nums[0] + previousSum

	return append([]int{s}, sum(nums[1:], s)...)
}

func runningSum(nums []int) []int {
	return sum(nums, 0)
}

func runningSum2(nums []int) []int {
	previousSum := 0
	result := make([]int, len(nums))

	for i, num := range nums {
		previousSum += num
		result[i] = previousSum
	}

	return result
}

func runningSum3(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}
