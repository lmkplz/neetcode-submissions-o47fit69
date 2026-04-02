func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []int{}

	for index, val := range(temperatures) {
		if index == 0 {
			stack = append(stack, index)
			continue
		}
		if val <= temperatures[stack[len(stack)-1]] {
			stack = append(stack, index)
			continue
		}
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < val {
			result[stack[len(stack)-1]] = index - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}

	return result
}
