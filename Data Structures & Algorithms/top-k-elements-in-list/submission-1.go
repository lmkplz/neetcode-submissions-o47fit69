
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	buckets := make([][]int, len(nums)+1)

	for num, count := range m {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)

	for i := len(buckets)-1; i > 0; i-- {
		bucket := buckets[i]

		for _, v := range bucket {
			result = append(result, v)

			if len(result) >= k {
				return result
			}
		}
	}

	return result
}
