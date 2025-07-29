package linked_list

import "github.com/y3933y3933/go-kata/implementations"

func cloneList(head *implementations.ListNode) *implementations.ListNode {
	if head == nil {
		return nil
	}

	dummy := &implementations.ListNode{Val: -1}
	tail := dummy

	for curr := head; curr != nil; curr = curr.Next {
		newNode := &implementations.ListNode{Val: curr.Val}
		tail.Next = newNode
		tail = tail.Next
	}
	return dummy.Next
}
