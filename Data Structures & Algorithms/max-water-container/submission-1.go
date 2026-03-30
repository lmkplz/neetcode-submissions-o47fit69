func maxArea(heights []int) int {
	l, r := 0, len(heights) - 1
	max := 0

	for l < r {
		val := (r - l) * min(heights[l], heights[r])
		if max < val {
			max = val
		}

		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}

	return max
}

