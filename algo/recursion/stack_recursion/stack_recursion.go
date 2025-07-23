package stackrecursion

func reverseStack(stack *[]int) {
	if len(*stack) == 0 {
		return
	}

	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	reverseStack(stack)

	insertAtBottom(stack, top)

}

func insertAtBottom(stack *[]int, elem int) {
	if len(*stack) == 0 {
		*stack = append(*stack, elem)
	} else {
		top := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]

		insertAtBottom(stack, elem)
		*stack = append(*stack, top)
	}
}
