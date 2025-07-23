package stackrecursion

import (
	"reflect"
	"testing"
)

func TestReverseStack(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "empty stack",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single element",
			input: []int{42},
			want:  []int{42},
		},
		{
			name:  "two elements",
			input: []int{1, 2},
			want:  []int{2, 1},
		},
		{
			name:  "multiple elements",
			input: []int{9, 5, 1, 2},
			want:  []int{2, 1, 5, 9},
		},
		{
			name:  "all same elements",
			input: []int{7, 7, 7},
			want:  []int{7, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := make([]int, len(tt.input))
			copy(stack, tt.input)
			reverseStack(&stack)
			if !reflect.DeepEqual(stack, tt.want) {
				t.Errorf("RecursivelyReverseStack(%v) = %v, want %v", tt.input, stack, tt.want)
			}
		})
	}
}
