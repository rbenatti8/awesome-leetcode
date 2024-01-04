package top_interview_questions_easy

func reverseString(s []byte) []byte {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}

	return s
}

func reverseString2(s []byte) []byte {
	length := len(s)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
