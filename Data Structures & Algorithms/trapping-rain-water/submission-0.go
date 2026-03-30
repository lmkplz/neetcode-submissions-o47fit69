func trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	for i := 0; i < len(height); i++ {
		if i == 0 {
			leftMax[i] = 0
		} else {
			if height[i - 1] > leftMax[i - 1] {
				leftMax[i] = height[i - 1]
			} else {
				leftMax[i] = leftMax[i - 1]
			}
		}
	}

	fmt.Println("leftMax: ", leftMax)

	for i := len(height) - 1; i >= 0; i-- {
		if i == len(height) - 1 {
			rightMax[i] = 0
		} else {
			if height[i + 1] > rightMax[i + 1] {
				rightMax[i] = height[i + 1]
			} else {
				rightMax[i] = rightMax[i + 1]
			}
		}
	}

	fmt.Println("rightMax: ", rightMax)

	sum := 0
	for i := 0; i < len(height); i++ {
		area := min(leftMax[i], rightMax[i]) - height[i]
		if area > 0 {
			sum += area
		}
	}

	return sum
}
