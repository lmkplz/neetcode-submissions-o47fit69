func minEatingSpeed(piles []int, h int) int {
	max, k := 0, 1

	for _, p := range piles {
		if max < p {
			max = p
		}
	}

	if len(piles) == h {
		return max
	}

	l, r := 1, max
	res := max

	for r >= l {
		count := 0
		k = (r + l) / 2

		for _, p := range piles {
			count += int(math.Ceil(float64(p) / float64(k)))
		}

		if count <= h {
			r = k - 1
			res = k
		} else if count > h {
			l = k + 1
		}
	}

	return res
}