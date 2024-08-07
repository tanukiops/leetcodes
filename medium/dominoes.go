package medium

const DominoLeft = 76
const DominoRight = 82
const DominoStable = 46

// it succeeds on all cases but runs out of time on leetcode ... needs to be optimized.
func PushDominoes(dominoes string) string {
	if len(dominoes) == 1 {
		return dominoes
	}
	result := dominoes
	for i := range dominoes {
		if dominoes[i] == DominoStable {
			nearestR := findNearestR(i, dominoes)
			nearestL := findNearestL(i, dominoes)
			switch {
			case nearestL < nearestR:
				result = replaceChar(result, i, DominoLeft)
			case nearestL > nearestR:
				result = replaceChar(result, i, DominoRight)
			}

		}
	}
	return result
}

func replaceChar(s string, index int, c rune) string {
	return s[:index] + string(c) + s[index+1:]
}

func findNearestR(index int, s string) int {
	count := 0
	for i := index; i >= 0; i-- {
		count++
		if s[i] == DominoLeft {
			return 999
		}
		if s[i] == DominoRight {
			return count
		}
	}
	return 999
}

func findNearestL(index int, s string) int {
	count := 0
	for i := index; i < len(s); i++ {
		count++
		if s[i] == DominoRight {
			return 999
		}
		if s[i] == DominoLeft {
			return count
		}
	}
	return 999
}
