package sortsquare

import (
	"slices"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{-4, -2, 1, 3}, []int{1, 4, 9, 16}},
		{[]int{-6, -5, -5}, []int{25, 25, 36}},
		{[]int{1, 2, 3}, []int{1, 4, 9}},
		{[]int{-1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{-3, -1, 0, 1, 2}, []int{0, 1, 1, 4, 9}},
	}

	for _, tt := range tests {
		got := SortedSquares(tt.input)
		if !slices.Equal(tt.expected, got) {
			t.Errorf("SortedSquares(%v) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
