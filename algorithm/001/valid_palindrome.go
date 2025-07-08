package valid_palindrome

import "strings"

// 請實作一個函式 IsPalindrome(s string) bool，
// 判斷輸入的字串 s 在以下處理後是否為一個回文（palindrome）。
// 回文的定義是：字串正著讀與反著讀皆相同。

func IsPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	lowerStr := strings.ToLower(s)

	for left < right {
		if !isAlphanumeric(lowerStr[left]) {
			left++
			continue
		}
		if !isAlphanumeric(lowerStr[right]) {
			right--
			continue
		}
		if lowerStr[left] != lowerStr[right] {
			return false
		}
		left++
		right--
	}

	return true

}

func isAlphanumeric(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}

	if b >= 'A' && b <= 'Z' {
		return true
	}

	if b >= 'a' && b <= 'z' {
		return true
	}

	return false
}
