package binary_search

func binarySearch(input []int, target int) bool {
	l := 0
	h := len(input)
	for l < h {
		m := l + (h-l)/2
		v := input[m]
		if v == target {
			return true
		} else if v < target {
			l = m + 1
		} else {
			h = m
		}

	}
	return false
}
