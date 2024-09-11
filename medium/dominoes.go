package medium

import "fmt"

const DominoLeft = 76
const DominoRight = 82
const DominoStable = 46

func PushDominoes(dominoes string) string {
	// divide dominoes per L
	var dominosets []string
	length := len(dominoes)
	for i, start := 0, 0; i < length; i++ {
		if dominoes[i] == 'L' || i == length-1 {
			dominosets = append(dominosets, dominoes[start:i+1])
			start = i + 1
		}
	}
	fmt.Println(dominosets)
	// for each divided domino set, run the pushing of dominos.
	dominoes = ""
	for i := 0; i < len(dominosets); i++ {
		dominoset := pushDominoesSubset(dominosets[i])
		dominoes += dominoset
	}
	// concat them together
	return dominoes
}
func pushDominoesSubset(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	leftActive := s[length-1] == 'L'
	for i, j := 0, length-1; j > -1; i++ {
		//check 1 ahead for both pointers if both L and R are triggered and the check ahead meets that becomes a .
		lIndex, rIndex := 0, 0
		replaceL, replaceR := false, false
		if i < length {
			for s[i] == '.' {
				i++
				if i == length {
					return s
				}
			}
			rIndex = i + 1
			if s[i] == 'R' {
				if i == length-1 {
					return s
				}
				if leftActive && rIndex == lIndex {
					return s
				}
				if s[rIndex] == 'L' {
					return s
				}
				if rIndex+1 >= length {
					return replaceChar(s, rIndex, 'R')
				}
				if s[rIndex+1] != 'L' {
					replaceR = true
				}
			}
		}
		if s[j] == 'L' {
			lIndex = j - 1
			if lIndex >= 1 {
				if s[lIndex-1] != 'R' && s[lIndex] != 'R' {
					replaceL = true
				}
			} else if lIndex == 0 {
				if s[lIndex] == '.' {
					replaceL = true
				}
			}

		}
		if replaceR {
			s = replaceChar(s, rIndex, 'R')
		}
		if replaceL {
			s = replaceChar(s, lIndex, 'L')
		}
		j--
	}
	return s
}

// it succeeds on all cases but runs out of time on leetcode ... needs to be optimized.
func PushDominoes2(dominoes string) string {
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
