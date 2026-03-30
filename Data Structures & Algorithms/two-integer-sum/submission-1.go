type IndexCount struct {
    index int
    count int
} 
func twoSum(nums []int, target int) []int {
    hash := make(map[int]IndexCount)
    for index, val := range nums {
        if entry, ok := hash[val]; !ok {
            hash[val] = IndexCount{ index: index, count: 1 }
        } else {
            hash[val] = IndexCount{ index: index, count: entry.count +1}
        }
    }
    for i, val := range nums {
        rest := target - val
        if rest != val && hash[rest].count > 0 {
            return []int{i, hash[rest].index}
        }

        if rest == val && hash[rest].count > 1 {
            return []int{i, hash[rest].index}
        }
    }

    return nil
}
