package two_sum

import "sort"

func twoSum(arr []int, target int) []int {
	sort.Ints(arr)
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return []int{arr[left], arr[right]}
		} else if sum > target {
			right--
		} else {
			left++
		}

	}
	return []int{}
}
