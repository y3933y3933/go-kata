package linked_list

type listNode struct {
	Val  int
	Next *listNode
}

/**
 * Definition for singly-linked list.
 * type listNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *listNode) *listNode {
	lenA := getListLength(headA)
	lenB := getListLength(headB)

	diff := diffInt(lenA, lenB)
	currA := headA
	currB := headB

	for i := 0; i < diff; i++ {
		if lenA < lenB {
			currB = currB.Next
		} else {
			currA = currA.Next
		}
	}

	for currA != nil && currB != nil {
		if currA == currB {
			return currA
		}
		currA = currA.Next
		currB = currB.Next
	}

	return nil

}

func getListLength(head *listNode) int {
	var length int
	for current := head; current != nil; current = current.Next {
		length++
	}

	return length
}

func diffInt(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
