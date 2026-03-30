func simplify(str string) string {
    hash := make(map[rune]int)

    for _, val := range str {
        hash[val]++
    }

    result := ""

    for i := 'a'; i <= 'z'; i++ {
        if count, ok := hash[i]; ok {
            result += string(i) + strconv.Itoa(count)
        }
    }

    return result
}

func groupAnagrams(strs []string) [][]string {
    result := [][]string{}

    hash := make(map[string][]int)

    for i, str := range strs {
        sStr := simplify(str)
        hash[sStr] = append(hash[sStr], i)
    }

    for _, arr := range hash {
        anagram := []string{}
        for _, index := range arr {
            anagram = append(anagram, strs[index])
        }
        result = append(result, anagram)
    }

    return result
}
