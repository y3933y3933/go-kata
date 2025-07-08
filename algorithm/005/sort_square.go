package sortsquare

// 請實作一個函式 SortedSquares(arr []int) []int，接收一個「非遞減排序（從小到大）」的整數陣列 arr，
// 並回傳一個新陣列，其中包含 每個元素平方後的結果，仍然是非遞減排序的順序。

func SortedSquares(arr []int) []int {
	left := 0
	right := len(arr) - 1
	var result []int

	for left <= right {
		absLeft := abs(arr[left])
		absRight := abs(arr[right])

		if absLeft > absRight {
			result = append(result, arr[left]*arr[left])
			left++
		} else {
			result = append(result, arr[right]*arr[right])
			right--
		}

	}

	reverse(result)
	return result
}

func reverse(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
