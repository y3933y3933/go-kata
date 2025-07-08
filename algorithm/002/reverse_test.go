package reverse

import (
	"slices"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input []rune
		want  []rune
	}{
		{input: []rune{'a', 'e', 'i', 'o', 'u'}, want: []rune{'u', 'o', 'i', 'e', 'a'}},
		{input: []rune{'a', 'b', 'c', 'd', 'e'}, want: []rune{'e', 'd', 'c', 'b', 'a'}},
		{input: []rune{}, want: []rune{}},
		{input: []rune{'a'}, want: []rune{'a'}},
		{input: []rune{'A', 'B'}, want: []rune{'B', 'A'}},
	}

	for _, tt := range tests {
		origin := make([]rune, len(tt.input))
		copy(origin, tt.input)

		Reverse(tt.input)
		if !slices.Equal(tt.input, tt.want) {
			t.Errorf("Reverse(%v) = %v; want %v", origin, tt.input, tt.want)
		}
	}
}
