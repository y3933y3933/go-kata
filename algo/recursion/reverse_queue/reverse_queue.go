package reversequeue

func reverseQueue(queue *[]int) {
	if len(*queue) == 0 {
		return
	}

	elem := (*queue)[0]
	*queue = (*queue)[1:]

	reverseQueue(queue)

	*queue = append(*queue, elem)
}
