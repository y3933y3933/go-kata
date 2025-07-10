package reversewordorder

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "This is a    string",
			expected: "string a is This",
		},
		{
			input:    "   fizz buzz   ",
			expected: "buzz fizz",
		},
		{
			input:    "random",
			expected: "random",
		},
		{
			input:    "   one   two   three   ",
			expected: "three two one",
		},
		{
			input:    "    ",
			expected: "",
		},
		{
			input:    " multiple     spaces    in   between ",
			expected: "between in spaces multiple",
		},
	}

	for _, tt := range tests {
		got := ReverseWords(tt.input)
		if got != tt.expected {
			t.Errorf("ReverseWords(%q) = %q; want %q", tt.input, got, tt.expected)
		}
	}
}
