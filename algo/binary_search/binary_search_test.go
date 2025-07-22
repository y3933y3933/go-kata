package binary_search

import "testing"

var testData = []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

func Test_binarySearch(t *testing.T) {
	test := []struct {
		input    []int
		target   int
		expected bool
	}{
		{testData, 1, true},
		{testData, 5, false},
		{testData, 99, true}}
	for _, tt := range test {
		got := binarySearch(tt.input, tt.target)
		if got != tt.expected {
			t.Errorf("linearSearch(%v, %d) = %v; want %v", tt.input, tt.target, got, tt.expected)
		}
	}
}
