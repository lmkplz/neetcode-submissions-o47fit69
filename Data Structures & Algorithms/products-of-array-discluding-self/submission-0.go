func productExceptSelf(nums []int) []int {
	size := len(nums)
	prefix := make([]int, size)
	suffix := make([]int, size)
	result := make([]int, size)

	for index := 0; index < size; index++ {
		if index == 0 {
			prefix[index] = 1
		} else {
			prefix[index] = prefix[index-1] * nums[index-1]
		}
	}

	for index := size - 1; index >= 0; index-- {
		if index == size - 1 {
			suffix[index] = 1
		} else {
			suffix[index] = suffix[index+1] * nums[index+1]
		}
	}

	for index := 0; index < size; index++ {
		result[index] = prefix[index] * suffix[index]
	}

	return result
}
