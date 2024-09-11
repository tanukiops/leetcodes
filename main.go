package main

import (
	"fmt"

	"github.com/tanukiops/leetcode/easy"
)

func main() {
	// result := easy.KthDistinct([]string{"a", "a", "aaa", "b", "b", "bb"}, 1)
	// fmt.Println(result)
	list := []int{1, 2, 9, 4, 5, 1003, 7, 8, 9, 50, 100, 110, 400, 700, 889, 2923, 23999, 400000}
	result := easy.TwoSum(list, 3)
	fmt.Println(result)

}

// func binsearch(list []int, target int, start int, end int) (int, bool) {
// 	fmt.Println("Searching")
// 	i := (start + end) / 2
// 	if list[i] == target {
// 		return i, true
// 	}
// 	if len(list[start:end]) == 1 {
// 		return 0, false
// 	} else if list[i] < target {
// 		return binsearch(list, target, i+1, end)
// 	} else if list[i] > target {
// 		return binsearch(list, target, start, i-1)
// 	}
// 	return 0, false
// }
// func binsearch2(list []int, target int) (int, bool) {
// 	fmt.Println("Searching")
// 	i := (len(list) / 2)
// 	switch {
// 	case len(list) == 0:
// 		return -1, false
// 	case list[i] > target:
// 		return binsearch2(list[:i], target)
// 	case list[i] < target:
// 		result, found := binsearch2(list[i:], target)
// 		if found {
// 			return i + result, true
// 		}
// 	default:
// 		return i, true
// 	}
// 	return -1, false
// }
