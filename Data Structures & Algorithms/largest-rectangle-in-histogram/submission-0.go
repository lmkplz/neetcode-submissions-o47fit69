func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([][2]int, n)
	max := 0

	for i := 0; i < n; i++ {
		cur := [2]int{i, heights[i]}
		for len(stack) > 0 && stack[len(stack) - 1][1] > cur[1] {
			size := stack[len(stack)-1][1] * (i - stack[len(stack)-1][0])
			if max < size {
				max = size
			}

			cur[0] = stack[len(stack)-1][0]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, cur)
	}

	for len(stack) > 0 {
		size := stack[len(stack) - 1][1] * (n - stack[len(stack)-1][0])
		if max < size {
			max = size
		}
		stack = stack[:len(stack)-1]
	}

	return max
}
