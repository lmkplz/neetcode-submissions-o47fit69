func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		fix := nums[i]

		if fix > 0 {
			break;
		}

		if i > 0 && fix == nums[i-1] {
			continue
		}

		r := len(nums) - 1
		l := i + 1

		for l < r {
			sum := fix + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{fix, nums[l], nums[r]})
				l++
				r--

				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
			if sum > 0 {
				r--
			}
			if sum < 0 {
				l++
			}
		}
	}

	return res
}
