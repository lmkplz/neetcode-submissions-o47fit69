func twoSum(numbers []int, target int) []int {
	hash := make(map[int]int)

	for index, num := range numbers {
		rest := target - num
		if val, exists := hash[rest]; exists {
			return []int{val, index + 1}
		}

		hash[num] = index + 1
	}

	return []int{}
}
