package medium

func LenghtOfLongestSubstring(s string) int {
	// according to the constrains, lenght can be zero
	if len(s) == 0 {
		return getLongestLength(s, 0)
	}
	return getLongestLength(s, 1)
}

// Used recursion to solve this problem, this is not the most efficient way
func getLongestLength(s string, length int) int {
	if len(s) > 1 && len(s) > length {
		letterMap := make(map[rune]int)
		for _, c := range s {
			letterMap[c]++
			if letterMap[c] > 1 {
				if len(letterMap) > length {
					return getLongestLength(s[1:], len(letterMap))
				} else {
					return getLongestLength(s[1:], length)
				}
			}
		}
		if len(letterMap) > length {
			return len(letterMap)
		}
	}
	return length
}
