package main

import (
	"fmt"

	"github.com/tanukiops/leetcode/medium"
)

func main() {
	// medium.LongestPalindrome("cbbd")
	node1 := &medium.ListNode{
		Val:  1,
		Next: nil,
	}
	node2 := &medium.ListNode{
		Val:  2,
		Next: nil,
	}
	node3 := &medium.ListNode{
		Val:  3,
		Next: nil,
	}
	node4 := &medium.ListNode{
		Val:  4,
		Next: nil,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	medium.PrintNodes(node1)
	fmt.Println("--- swap pairs ---")
	medium.SwapPairs(nil)
	medium.PrintNodes(node2)
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
