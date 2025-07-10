package reversewords

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"This is a string", "sihT si a gnirts"},
		{"I  love   coding", "I  evol   gnidoc"},
		{"random", "modnar"},
		{"   ", "   "},
		{"a", "a"},
	}

	for _, tt := range tests {
		output := ReverseWords(tt.input)

		if output != tt.expect {
			t.Errorf("ReverseWords(%q) = %q; want %q", tt.input, output, tt.expect)
		}
	}
}
