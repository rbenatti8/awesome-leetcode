package top_interview_questions_easy

func isAnagram(s string, t string) bool {
	var m, n = [26]int{}, [26]int{}

	const AsciiA = int('a')

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		m[int(s[i])-AsciiA]++
		n[int(t[i])-AsciiA]++
	}

	for i := 0; i < 26; i++ {
		if m[i] != n[i] {
			return false
		}
	}

	return true
}

func isAnagram2(s string, t string) bool {
	var m = [26]int{}

	const AsciiA = int('a')

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		m[int(s[i])-AsciiA]++
		m[int(t[i])-AsciiA]--
	}

	for i := 0; i < 26; i++ {
		if m[i] != 0 {
			return false
		}
	}

	return true
}
