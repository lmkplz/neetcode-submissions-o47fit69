func productExceptSelf(nums []int) []int {
	size := len(nums)
	result := make([]int, size)

	// set prefix value to result
	prefix := 1
	for index := 0; index < size; index++ {
		result[index] = prefix
		prefix *= nums[index]
	}

	fmt.Println(result)

	// calculate suffix value and store it in a variable. 
	// Then multiply suffix value and prefix value from array
	suffix := 1

	for index := size - 1; index >= 0; index-- {
		result[index] = result[index] * suffix
		suffix *= nums[index]
	}

	return result
}
