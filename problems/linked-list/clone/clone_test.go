package linked_list

import (
	"slices"
	"testing"

	"github.com/y3933y3933/go-kata/implementations"
)

func TestCloneList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{0, 1, 2}, expected: []int{0, 1, 2}},
		{input: []int{}, expected: []int{}},
		{input: []int{42}, expected: []int{42}},
	}

	for _, tt := range tests {
		original := buildList(tt.input)
		cloned := cloneList(original)

		got := listToSlice(cloned)

		if !slices.Equal(got, tt.expected) {
			t.Errorf("cloneList(%v) = %v; want %v", tt.input, got, tt.expected)
		}

	}
}

func buildList(vals []int) *implementations.ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &implementations.ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &implementations.ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func listToSlice(head *implementations.ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
