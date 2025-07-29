package clone

import (
	"slices"
	"testing"

	linked_list "github.com/y3933y3933/go-kata/problems/linked-list"
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
		original := linked_list.BuildList(tt.input)
		cloned := cloneList(original)

		got := linked_list.ListToSlice(cloned)

		if !slices.Equal(got, tt.expected) {
			t.Errorf("cloneList(%v) = %v; want %v", tt.input, got, tt.expected)
		}

	}
}
