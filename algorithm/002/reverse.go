package reverse

// 請實作一個函式 Reverse(arr []rune)，**就地（in-place）**反轉一個字元陣列 arr。
// 你必須直接修改輸入陣列的內容（不能建立新的陣列），並且保證空間複雜度為 O(1)。

func Reverse(arr []rune) {
	left := 0
	right := len(arr) - 1
	for left < right {
		tmp := arr[left]
		arr[left] = arr[right]
		arr[right] = tmp

		left++
		right--
	}
}
