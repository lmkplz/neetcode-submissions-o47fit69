func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1

	for r > l {
		m := (r + l) / 2 
		if nums[m] < nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}

	pivot := l
	l, r = 0, n-1

	if target >= nums[pivot] && target <= nums[r] {
		l = pivot
	} else {
		r = pivot - 1
	}

	fmt.Println(l, r)	

	for r >= l {
		m := (r + l)/2
		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}
