package swap

// 請實作一個 Swap(a, b *int) 函式，能夠交換兩個整數變數的值。

func Swap(a, b *int) {
	tmp := *b
	*b = *a
	*a = tmp
}
