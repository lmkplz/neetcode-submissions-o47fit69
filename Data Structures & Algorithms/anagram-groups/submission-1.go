func groupAnagrams(strs []string) [][]string {
    hash := make(map[string][]string)

    for i := range strs {
        sRunes := []byte(strs[i])
        sort.Slice(sRunes, func(i, j int) bool {
            return sRunes[i] < sRunes[j]
        })

        anagram := string(sRunes)
        hash[anagram] = append(hash[anagram], strs[i])
    }

    result := make([][]string, 0, len(hash))
    for _, v := range hash {
        result = append(result, v)
    }

    return result
}
