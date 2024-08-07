package medium

import "slices"

func MinimumPushes(word string) int {
	distinctLetterCount, letters := getDistinctLetters(word)
	slices.Sort(letters)
	result := 0
	clicks := 1
	for i := 0; i < distinctLetterCount; i++ {
		if i > 0 && i%8 == 0 {
			clicks++
		}
		result += clicks * letters[distinctLetterCount-i-1]
	}
	return result
}

func getDistinctLetters(word string) (int, []int) {
	letterMap := make(map[rune]int)
	for _, c := range word {
		letterMap[c]++
	}
	keypresses := make([]int, len(letterMap))
	i := 0
	for _, v := range letterMap {
		keypresses[i] = v
		i++
	}
	return len(letterMap), keypresses
}
