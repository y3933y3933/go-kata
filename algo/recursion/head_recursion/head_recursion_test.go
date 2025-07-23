package headrecursion

import (
	"slices"
	"testing"
)

func TestRecursivelyPrintNumbersFrom1ToN(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "n=1",
			n:    1,
			want: []int{1},
		},
		{
			name: "n=5",
			n:    5,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "n=0",
			n:    0,
			want: []int{},
		},
		{
			name: "n=10",
			n:    10,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := recursivelyPrintNumbersFrom1ToN(tt.n)
			if !slices.Equal(got, tt.want) {
				t.Errorf("recursivelyPrintNumbersFrom1ToN(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}

}
