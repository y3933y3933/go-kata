package reversewords

import (
	"strings"
)

// 請實作一個函式 ReverseWords(s string)，將輸入的字元陣列 chars 中的每個單字反轉，但保留原本單字的順序與空格。

func ReverseWords(s string) string {
	words := strings.Split(s, " ")

	for i, word := range words {

		chars := []rune(word)
		reverse(chars, 0, len(word)-1)

		words[i] = string(chars)
	}

	return strings.Join(words, " ")
}

func reverse(runes []rune, left, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}
