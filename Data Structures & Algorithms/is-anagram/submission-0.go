func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    convertSortStr := func(str string) string {
        hash := make(map[rune]int)

        for _, ch := range str {
            hash[ch]++
        }

        result := ""
        for i := 'a'; i <= 'z'; i++ {
            if val, ok := hash[i]; ok {
                result += string(i) + strconv.Itoa(val)
            }
        }

        return result
    }

    first := convertSortStr(s)
    second := convertSortStr(t)

    return first == second
}
