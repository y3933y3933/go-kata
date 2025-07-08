package reversekgroup

// 請實作一個函式 ReverseStr(s string, k int) string，按照以下規則處理字串：
// 對於從字串起始開始的每個 2k 區間：
// 1. 反轉前 k 個字元
// 2. 保留後 k 個字元不變
// 如果剩下的字元小於 k，全部反轉
// 如果剩下的字元在 k 到 2k 之間，只反轉前 k 個，其餘保留

func ReverseStr(s string, k int) string {
	chars := []rune(s)
	for i := 0; i < len(s); i += 2 * k {
		endIndex := min(i+k-1, len(chars)-1)
		reverse(chars, i, endIndex)

	}
	return string(chars)
}

func reverse(chars []rune, left, right int) {
	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--

	}
}
