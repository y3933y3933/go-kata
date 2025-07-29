package clone

import linked_list "github.com/y3933y3933/go-kata/problems/linked-list"

func cloneList(head *linked_list.ListNode) *linked_list.ListNode {
	if head == nil {
		return nil
	}

	dummy := &linked_list.ListNode{Val: -1}
	tail := dummy

	for curr := head; curr != nil; curr = curr.Next {
		newNode := &linked_list.ListNode{Val: curr.Val}
		tail.Next = newNode
		tail = tail.Next
	}
	return dummy.Next
}
