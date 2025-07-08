package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"a man nam a", true},
		{"race car rac ecar", true},
		{"12321", true},
		{"1a2", false},
		{"Able , was I saw eLba", true},
		{"", true},
		{"A", true},
		{"ab", false},
	}

	for _, tt := range tests {
		got := IsPalindrome(tt.input)
		if got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
