package fibonaccinumber

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{10, 55},
	}

	for _, tt := range tests {
		got := fibonacci(tt.input)
		if got != tt.expected {
			t.Errorf("Fibonacci(%d) = %d; want %d", tt.input, got, tt.expected)
		}
	}
}
