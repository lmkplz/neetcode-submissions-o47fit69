func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	
	l := 0
	defaultHash := make(map[byte]int)

	for l = 0; l < len(s1); l++ {
		defaultHash[s1[l]]++
	}

	for l = 0; l < len(s2) - len(s1) + 1; l++ {
		tmp := s2[l:l+len(s1)]
		hash := make(map[byte]int)

		for i := 0; i < len(tmp); i++ {
			hash[tmp[i]]++
		}

		if len(hash) != len(defaultHash) {
			continue
		}

		numOfChars := 0

		for key, val := range hash {
			if val == defaultHash[key] {
				numOfChars++
			}
		}
		if numOfChars == len(hash) {
			return true
		}
	}

	return false
}
