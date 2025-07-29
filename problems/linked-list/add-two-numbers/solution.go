package addtwonumber

import linked_list "github.com/y3933y3933/go-kata/problems/linked-list"

func addTwoNumbers(l1 *linked_list.ListNode, l2 *linked_list.ListNode) *linked_list.ListNode {
	dummy := &linked_list.ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curr.Next = &linked_list.ListNode{Val: sum % 10}
		curr = curr.Next
		carry = sum / 10
	}

	return dummy.Next
}
