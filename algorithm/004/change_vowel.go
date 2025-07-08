package changevowel

// 請實作一個函式 ReverseVowels(s string) string，
// 將字串中的 **母音字元（vowels）**反轉，其它字元保持原位不動。

func ReverseVowels(s string) string {
	chars := []byte(s)
	left := 0
	right := len(s) - 1
	for left < right {
		if !isVowel(chars[left]) {
			left++
			continue
		}

		if !isVowel(chars[right]) {
			right--
			continue
		}

		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--

	}
	return string(chars)
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}

}
