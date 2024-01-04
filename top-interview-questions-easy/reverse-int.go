package top_interview_questions_easy

func reverseInt(x int) int {
	var res int

	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}

	if res > 2147483647 || res < -2147483647 {
		return 0
	}

	return res
}
