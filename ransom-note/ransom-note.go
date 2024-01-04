package ransom_note

// O(n+m) time complexity
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, r := range magazine {
		m[r]++
	}

	for _, r := range ransomNote {
		if m[r] == 0 {
			return false
		}

		m[r]--
	}

	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	m := make(map[rune]int, 26)

	s := magazine + "-" + ransomNote
	shouldDecrement := false
	for _, r := range s {
		if r == '-' {
			shouldDecrement = true
			continue
		}

		if shouldDecrement {
			if m[r] == 0 {
				return false
			}

			m[r]--
		} else {
			m[r]++
		}

	}

	return true
}

// In ASCII, the value of 'a' is 97 and 'z' is 122. When you subtract 'a' from a character, for instance, 'c' (whose ASCII value is 99), c - 'a' results in 99 - 97 = 2.
// This yields the index 2 in the array, which corresponds to the letter 'c' in the English alphabet (0-indexed).
func canConstruct3(ransomNote string, magazine string) bool {
	var res [26]int
	var res2 [26]int

	for _, c := range magazine {
		res2[c-'a']++
	}

	for _, v := range ransomNote {
		res[v-'a']++

		if res[v-'a'] > res2[v-'a'] {
			return false
		}
	}

	return true
}

func canConstruct4(ransomNote string, magazine string) bool {
	var ransomNoteLetters [26]int
	var magazineLetters [26]int

	for i := 0; i < len(magazine); i++ {
		magazineLetters[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		index := ransomNote[i] - 'a'

		ransomNoteLetters[index]++

		if ransomNoteLetters[index] > magazineLetters[index] {
			return false
		}
	}

	return true
}
