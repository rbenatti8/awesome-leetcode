package top_interview_questions_easy

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sig := 1
	i := 0
	val := 0

	for i != len(s) {
		if s[i] == ' ' {
			i++
			continue
		}

		if s[i] == '-' {
			sig = -1
			i++
			continue
		}

		if isNumericRune(s[i]) == false {
			return sig * val
		}

		val = val*10 + int(s[i]) - '0'

		if val > 2147483647 || val < -2147483648 {
			if sig == 1 {
				return 2147483647
			}

			return -2147483648
		}

		i++
	}

	return sig * val
}

func myAtoi2(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1 // -1 negative
	index := 0
	val := 0

	for len(s) > index && s[index] == ' ' {
		index++
	}

	if index >= len(s) {
		return 0
	}

	if v := s[index]; v == '-' || v == '+' {
		if v == '-' {
			sign = -1
		}

		index++
	}

	for {
		if index == len(s) {
			break
		}

		if isNumericRune(s[index]) == false {
			return sign * val
		}

		val = val*10 + int(s[index]) - '0'

		if val > 2147483647 || val < -2147483648 {
			if sign == 1 {
				return 2147483647
			}

			return -2147483648
		}

		index++
	}

	return sign * val
}

func isNumericRune(x byte) bool {
	return x >= '0' && x <= '9'
}
