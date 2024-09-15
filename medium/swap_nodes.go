package medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		// last node in the list, just return iself
		return head
	}
	// if head.Next.Next == nil {
	// 	//last 2 nodes in the list, swap and return head
	// 	newHead := head.Next
	// 	newHead.Next = head
	// 	head.Next = nil
	// 	return newHead
	// }
	//swap head and head.next then go deeper on head.next
	node3 := head.Next.Next
	newHead := head.Next
	newHead.Next = head
	head.Next = SwapPairs(node3)
	return newHead
}
func PrintNodes(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Println(curr)
		curr = curr.Next
	}
}
