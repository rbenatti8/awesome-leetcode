package number_of_steps_to_reduce_a_number_to_zero

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}

	if num%2 == 0 {
		return 1 + numberOfSteps(num/2)
	} else {
		return 1 + numberOfSteps(num-1)
	}
}

func numberOfSteps2(num int) int {
	steps := 0

	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}

		steps++
	}

	return steps
}

func numberOfSteps3(num int) int {
	steps := 0

	for num > 0 {
		if num&1 == 0 {
			num = num >> 1
		} else {
			num = num ^ 1
		}

		steps++
	}

	return steps
}

func numberOfSteps4(num int) int {
	if num == 0 {
		return 0
	}

	step := 0

	for num > 0 {
		step += num&1 + 1
		num = num >> 1
	}

	return step - 1
}
