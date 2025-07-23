package headrecursion

// Example
// input: n = 10
// output: [1,2,3,4,5,6,7,8,9,10]
// 必須要用 recursion 處理
func recursivelyPrintNumbersFrom1ToN(n int) []int {
	result := make([]int, n)
	helper(result, n)
	return result
}

func helper(result []int, n int) {
	if n == 0 {
		return
	}
	helper(result, n-1)

	result[n-1] = n
}
