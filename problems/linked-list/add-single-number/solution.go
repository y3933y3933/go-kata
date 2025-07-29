package addsinglenumber

import (
	linked_list "github.com/y3933y3933/go-kata/problems/linked-list"
)

func addSingleNumber(head *linked_list.ListNode, digit int) *linked_list.ListNode {
	carry := digit

	curr := head
	var prev *linked_list.ListNode

	for curr != nil && carry > 0 {
		sum := carry + curr.Val
		curr.Val = sum % 10

		carry = sum / 10

		prev = curr
		curr = curr.Next
	}

	if prev != nil && carry > 0 {
		prev.Next = &linked_list.ListNode{Val: carry}
	}

	return head
}
