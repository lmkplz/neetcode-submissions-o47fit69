func lengthOfLongestSubstring(s string) int {
	l, res := 0, 0
	hash := make(map[byte]int)

	for r := 0; r < len(s); r++ {
		if val, exists := hash[s[r]]; exists {
			l = max(val + 1, l)
		}

		hash[s[r]] = r
		res = max(r - l + 1, res)
	}

	return res
}
