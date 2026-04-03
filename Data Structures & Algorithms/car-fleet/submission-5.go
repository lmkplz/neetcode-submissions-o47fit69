func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	pairs := make([][2]int, n)
	stack := []float32{}

	for i := 0; i < n; i++ {
		pairs[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	for _, pair := range pairs {
		time := float32(target - pair[0]) / float32(pair[1])
		stack = append(stack, time)
		if len(stack) >= 2 && stack[len(stack)-2] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}
