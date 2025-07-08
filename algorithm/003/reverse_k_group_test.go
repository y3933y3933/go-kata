package reversekgroup

import "testing"

func TestReverseStr(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		expect string
	}{
		{"abcdefghij", 2, "bacdfeghji"},
		{"dfgh", 5, "hgfd"},
		{"qwerty", 3, "ewqrty"},
		{"abcdefg", 4, "dcbaefg"},
		{"a", 1, "a"},
		{"", 3, ""},
	}

	for _, tt := range tests {
		got := ReverseStr(tt.s, tt.k)
		if got != tt.expect {
			t.Errorf("ReverseStr(%q, %d) = %q; want %q", tt.s, tt.k, got, tt.expect)
		}
	}
}
