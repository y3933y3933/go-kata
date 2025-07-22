package two_sum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   []int
	}{
		{
			arr:    []int{2, 8, 3, 6, 4},
			target: 7,
			want:   []int{3, 4},
		},
		{
			arr:    []int{2, -1, 5, -4, 3},
			target: 34,
			want:   []int{},
		},
		{
			arr:    []int{2},
			target: 2,
			want:   []int{},
		},
		{
			arr:    []int{-3, 1, 4, 2, -2},
			target: -1,
			want:   []int{-3, 2},
		},
		{
			arr:    []int{0, 1, 2, -2},
			target: 0,
			want:   []int{2, -2},
		},
		{
			arr:    []int{1, 2, 3, 4},
			target: 6,
			want:   []int{2, 4},
		},
		{
			arr:    []int{5, 5, 1},
			target: 10,
			want:   []int{5, 5},
		},
	}

	for _, tt := range tests {
		got := twoSum(tt.arr, tt.target)
		if len(got) == 0 && len(tt.want) == 0 {
			return
		}
		if !slices.Equal(tt.want, got) {
			t.Errorf("twoSum(%v, %d) = %v, want %v", tt.arr, tt.target, got, tt.want)
		}

	}
}
