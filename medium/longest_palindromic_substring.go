package medium

func LongestPalindrome(s string) string {
	length := len(s)
	for width := length; width > 0; width-- {
		slider := length - width
		for i := 0; i <= slider; i++ {
			if isPalindrome(s[i : width+i]) {
				return s[i : width+i]
			}
		}
		//decrease word width
	}
	return ""
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	i, j := 0, len(s)-1
	for {
		if s[i] != s[j] {
			//not a palindrome
			return false
		}
		if i == j {
			return true
		}
		if len(s)%2 == 0 {
			if i+1 == j {
				return true
			}
		}
		i++
		j--
	}
}
