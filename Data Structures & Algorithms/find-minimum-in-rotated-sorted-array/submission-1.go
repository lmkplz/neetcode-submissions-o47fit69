func findMin(nums []int) int {
	n, min := len(nums), nums[0]
	l, r := 0, n-1

	for r >= l {
		m := (r + l) / 2 
		if min > nums[m] {
			min = nums[m]
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return min
}
 