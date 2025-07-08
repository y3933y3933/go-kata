package changevowel

import "testing"

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"random", "rondam"},
		{"afegijoku", "ufogijeka"},
		{"bcdf", "bcdf"},
		{"aeiou", "uoiea"},
		{"", ""},
		{"hello", "holle"},
	}

	for _, tt := range tests {
		got := ReverseVowels(tt.input)
		if got != tt.expect {
			t.Errorf("ReverseVowels(%q) = %q; want %q", tt.input, got, tt.expect)
		}
	}
}
