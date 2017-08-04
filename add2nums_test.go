package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := ListNode{}
	cur := &root
	var carry int
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		curRes := getVal(l1) + getVal(l2) + carry
		cur.Next = &ListNode{curRes % 10, nil}
		cur = cur.Next
		carry = (curRes) / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return root.Next
}

func getVal(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}
