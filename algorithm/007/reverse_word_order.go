package reversewordorder

import (
	"strings"
)

// 給定一個字串 s，請撰寫一個函式 ReverseWords(s string) string，將字串中的單字順序反轉，
// 並使用一個空白字元作為分隔。輸入可能包含前導、後導空白或是多個空白，輸出字串中不應該有多餘空白。
func ReverseWords(s string) string {
	words := strings.Fields(s)
	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}
