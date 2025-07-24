package binomial

func binomial(n int, k int) int {
	if k == 0 || n == k {
		return 1
	}

	return binomial(n-1, k) + binomial(n-1, k-1)
}
