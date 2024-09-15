package medium

func SwapNodes(head *ListNode, k int) *ListNode {
	nodes := linkedListToArray(head)
	length := len(nodes)
	k2 := length - k
	temp := nodes[k-1]
	nodes[k-1] = nodes[k2]
	nodes[k2] = temp
	return arrayToLinkedList(nodes)
}

func linkedListToArray(head *ListNode) (nodes []*ListNode) {
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	return
}

func arrayToLinkedList(nodes []*ListNode) (head *ListNode) {
	for i := 0; i < len(nodes); i++ {
		nodes[i].Next = nil
		if i+1 < len(nodes) {
			nodes[i].Next = nodes[i+1]
		}
	}
	return nodes[0]
}
