func twoSum(numbers []int, target int) []int {
	hash := make(map[int]int)
	for k, v := range numbers {
		hash[v] = k + 1
	}

	for index, num := range numbers {
		rest := target - num
		if hash[rest] > 0 {
			return []int{index + 1, hash[rest]}
		}
	}

	return []int{}
}
