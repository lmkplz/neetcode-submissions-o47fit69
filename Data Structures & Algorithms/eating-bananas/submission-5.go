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

	l, r := 0, max
	res := [2]int{len(piles), max}

	for r - 1 > l {
		count := 0
		k = (r + l) / 2

		for _, p := range piles {
			count += int(math.Ceil(float64(p) / float64(k)))
		}

		if count <= h {
			r = k
			if res[0] < count {
				res[0] = count
				res[1] = k
			} else if res[0] == count {
				if res[1] > k {
					res[1] = k
				}
			}
		} else if count > h {
			l = k
		}
	}

	return res[1]
}