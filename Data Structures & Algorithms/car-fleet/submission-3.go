func carFleet(target int, position []int, speed []int) int {
	sortedTime := make([]float32, len(position))
	stack := []float32{}

	for i:=0; i<len(position); i++ {
		for j:=i+1; j<len(position); j++ {
			if position[i] < position[j] {
				tmp := position[j]
				position[j] = position[i]
				position[i] = tmp

				tmp = speed[j]
				speed[j] = speed[i]
				speed[i] = tmp
			}
		}
		sortedTime[i] = (float32(target) - float32(position[i])) / float32(speed[i])
	}

	for k, t := range sortedTime {
		if k == 0 {
			stack = append(stack, t)
			continue
		}
		if len(stack) > 0 && stack[len(stack)-1] < t {
			stack = append(stack, t)
		}
	}

	fmt.Println(stack)

	return len(stack)
}
