func characterReplacement(s string, k int) int {
	l, maxf, res := 0, 0, 0
	count := make(map[byte]int)

	for r := 0; r < len(s); r++ {
		count[s[r]]++

		if count[s[r]] > maxf {
			maxf = count[s[r]]
		}

		for (r - l + 1) - maxf > k {
			count[s[l]]--
			l++
		}

		if r - l + 1 > res {
			res = r - l + 1
		}
	}

	return res
}
