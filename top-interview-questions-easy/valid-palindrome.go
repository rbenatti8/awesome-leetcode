package top_interview_questions_easy

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")

	if len(s) == 0 {
		return true
	}

	for i, j := 0, len(s)-1; i < j; {
		if !unicode.IsLetter(rune(s[i])) {
			s = s[:i] + s[i+1:]
		}

		if !unicode.IsLetter(rune(s[j])) {
			s = s[:j] + s[j+1:]
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}

func isPalindrome2(s string) bool {
	chars := []rune(s)

	for i, j := 0, len(chars)-1; i < j; {
		if !unicode.IsLetter(chars[i]) && !unicode.IsNumber(chars[i]) {
			i++
			continue
		}

		if !unicode.IsLetter(chars[j]) && !unicode.IsNumber(chars[j]) {
			j--
			continue
		}

		if unicode.ToLower(chars[i]) != unicode.ToLower(chars[j]) {
			return false
		}

		i++
		j--
	}

	return true
}

func isPalindrome3(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}
	s = strings.Map(f, s)

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i = i + 1
		j = j - 1
	}

	return true
}

func isPalindrome4(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !(unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i]))) {
			i++
		}
		for i < j && !(unicode.IsLetter(rune(s[j])) || unicode.IsDigit(rune(s[j]))) {
			j--
		}
		if i == j {
			break
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
	}
	return true
}

func isPalindrome5(s string) bool {
	for i, j := 0, len(s)-1; i < j; {

		if !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i++
			continue
		}

		if !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j--
			continue
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}
