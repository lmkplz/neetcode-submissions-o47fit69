func hasDuplicate(nums []int) bool {
    count := make(map[int]int)

    for _, n := range nums {
        count[n]++
        if count[n] > 1 {
            return true
        }
    }
    return false
}
