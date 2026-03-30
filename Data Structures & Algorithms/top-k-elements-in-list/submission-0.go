type Freq struct {
	Num int
	Count int
}
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	keys := make([]Freq, 0, len(m))

	for k, v := range m {
		keys = append(keys, Freq{ Num: k, Count: v })
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Count > keys[j].Count
	})

	fmt.Println(keys)

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, keys[i].Num)
	}

	return result
}
