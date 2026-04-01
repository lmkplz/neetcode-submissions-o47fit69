func maxSlidingWindow(nums []int, k int) []int {
    l, r := 0, 0
	deque := []int{}
	result := []int{}

	for r < len(nums) {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[r] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, r)

		if deque[0] < l {
			deque = deque[1:]
		}

		if r+1 >= k {
			result = append(result, nums[deque[0]])
			l++
		}
		r++
	}

	return result
}
