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
	node5 := &medium.ListNode{
		Val:  5,
		Next: nil,
	}
	// node6 := &medium.ListNode{
	// 	Val:  6,
	// 	Next: nil,
	// }
	// node7 := &medium.ListNode{
	// 	Val:  7,
	// 	Next: nil,
	// }
	// node8 := &medium.ListNode{
	// 	Val:  8,
	// 	Next: nil,
	// }
	// node9 := &medium.ListNode{
	// 	Val:  9,
	// 	Next: nil,
	// }
	// node10 := &medium.ListNode{
	// 	Val:  10,
	// 	Next: nil,
	// }
	// node11 := &medium.ListNode{
	// 	Val:  11,
	// 	Next: nil,
	// }
	node1.Next = node2
	node2.Next = nil
	node3.Next = node4
	node4.Next = node5
	// node5.Next = node6
	// node6.Next = node7
	// node7.Next = node8
	// node8.Next = node9
	// node9.Next = node10
	// node10.Next = node11

	// medium.PrintNodes(node1)
	// fmt.Println("--- swap  ---")
	medium.SwapNodes(node1, 1)
	fmt.Println("--- swapped  ---")
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
