func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	size, matches := len(s1), 0
	s1Count := make([]int, 26)
	s2Count := make([]int, 26)

	fmt.Println(size)

	for i := 0; i < size; i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}
	if matches == 26 {
		return true
	}

	for l := 0; l < len(s2) - size; l++ {
		matches = 0
		
		if s2[l] != s2[l+size] {
			s2Count[s2[l]-'a']--
			s2Count[s2[l+size]-'a']++

			fmt.Println(s1Count, s2Count)
		}

		for i := 0; i < 26; i++ {
			if s1Count[i] == s2Count[i] {
				matches++
			}
		}
		if matches == 26 {
			return true
		}
	}

	return false
}
