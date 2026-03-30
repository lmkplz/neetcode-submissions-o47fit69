func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	hash := make(map[int]struct{})

	for _, num := range nums {
		hash[num] = struct{}{}
	}
	
	result := 1

	for num := range hash {
		if _, found := hash[num-1]; !found {
			cur := num + 1
			longest := 1 
			for {
				if _, exists := hash[cur]; !exists {
					break	
				}
				cur++
				longest++
			}

			if result < longest {
				result = longest
			}
		}
	}

	return result
}
