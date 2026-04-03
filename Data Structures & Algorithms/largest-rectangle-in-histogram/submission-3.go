func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([][2]int, n)
	max := 0

	for i := 0; i < n; i++ {
		cur := [2]int{i, heights[i]}
		for len(stack) > 0 && stack[len(stack) - 1][1] > cur[1] {
			area := stack[len(stack)-1][1] * (i - stack[len(stack)-1][0])
			if max < area {
				max = area
			}

			cur[0] = stack[len(stack)-1][0]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, cur)
	}

	for _, pair := range stack {
		area := pair[1] * (n - pair[0])
		if max < area {
			max = area
		}
	}

	return max
}
