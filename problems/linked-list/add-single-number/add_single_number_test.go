package addsinglenumber

import (
	"slices"
	"testing"

	linked_list "github.com/y3933y3933/go-kata/problems/linked-list"
)

func TestAddSingleNumber(t *testing.T) {
	tests := []struct {
		input  []int
		digit  int
		expect []int
	}{
		{

			input:  []int{2, 4, 3},
			digit:  9,
			expect: []int{1, 5, 3},
		},
		{

			input:  []int{7, 9, 9},
			digit:  4,
			expect: []int{1, 0, 0, 1},
		},
		{

			input:  []int{1, 2, 3},
			digit:  4,
			expect: []int{5, 2, 3},
		},
		{

			input:  []int{9, 9, 9},
			digit:  1,
			expect: []int{0, 0, 0, 1},
		},
		{

			input:  []int{5},
			digit:  2,
			expect: []int{7},
		},
		{

			input:  []int{9},
			digit:  2,
			expect: []int{1, 1},
		},
	}

	for _, tt := range tests {
		head := linked_list.BuildList(tt.input)
		got := addSingleNumber(head, tt.digit)
		gotSlice := linked_list.ListToSlice(got)

		if !slices.Equal(gotSlice, tt.expect) {
			t.Errorf("addSingleNumber(%v, %d) = %v, want %v", tt.input, tt.digit, gotSlice, tt.expect)
		}

	}
}
