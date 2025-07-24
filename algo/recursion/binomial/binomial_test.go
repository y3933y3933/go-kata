package binomial

import "testing"

func TestBinomialCoefficient(t *testing.T) {
	tests := []struct {
		n, k     int
		expected int
	}{
		{5, 3, 10},
		{10, 4, 210},
		{6, 2, 15},
		{7, 0, 1},
		{7, 7, 1},
		{8, 1, 8},
		{8, 7, 8},
		{9, 3, 84},
		{12, 6, 924},
	}

	for _, tt := range tests {
		got := binomial(tt.n, tt.k)
		if got != tt.expected {
			t.Errorf("BinomialCoefficient(%d, %d) = %d; want %d", tt.n, tt.k, got, tt.expected)
		}
	}
}
