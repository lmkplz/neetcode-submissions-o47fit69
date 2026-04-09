func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	idx, idx1, idx2 := 0, 0, 0
	m, n := len(nums1), len(nums2)
	mid1, mid2  := 0, 0

	if (m+n)%2 == 1 {
		mid1 = (m+n)/2
		mid2 = mid1
	} else {
		mid1 = (m+n)/2
		mid2 = mid1 - 1
	}
	
	nums := []int{}

	if m == 0 {
		nums = nums2
	} else if n == 0 {
		nums = nums1
	} else {
		for idx <= mid1 {
			if nums1[idx1] < nums2[idx2] {
				nums = append(nums, nums1[idx1])
				idx1++
			} else {
				nums = append(nums, nums2[idx2])
				idx2++
			}
			idx++

			if idx1 == m {
				nums = append(nums, nums2[idx2:]...)
				break
			}
			if idx2 == n {
				nums = append(nums, nums1[idx1:]...)
				break
			}
		}
	}

	return float64(nums[mid1] + nums[mid2]) / 2
}
