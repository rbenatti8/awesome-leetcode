package top_interview_questions_easy

func firstUniqChar(s string) int {
	r := make(map[rune]int)

	for _, c := range s {
		r[c]++
	}

	for i, c := range s {
		if r[c] == 1 {
			return i
		}
	}

	return -1
}

func firstUniqChar2(s string) int {
	flags := make([]int, 26)
	for i := range flags {
		flags[i] = -1
	}
	slen := len(s)

	for i, ch := range s {
		idx := byte(ch - 'a')
		if flags[idx] == -1 {
			flags[idx] = i
		} else {
			flags[idx] = slen
		}
	}

	min := slen
	for i := range flags {
		if flags[i] >= 0 && flags[i] < len(s) && flags[i] < min {
			min = flags[i]
		}
	}

	if min == slen {
		return -1
	}
	return min
}

func firstUniqChar3(s string) int {
	var r = [26]int{}

	for _, c := range s {
		r[c-'a']++
	}

	for i, c := range s {
		if r[c-'a'] == 1 {
			return i
		}
	}

	return -1
}
