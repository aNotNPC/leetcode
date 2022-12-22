package medianoftwosortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	allLen := len1 + len2
	nums := make([]int, 0, allLen)

	i, j := 0, 0
	for {
		if i >= len1 || j >= len2 {
			break
		}
		if nums1[i] <= nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}

	}

	if i < len1 {
		nums = append(nums, nums1[i:]...)
	} else {
		nums = append(nums, nums2[j:]...)
	}

	mid := allLen / 2
	if allLen%2 == 1 {
		return float64(nums[mid])
	} else {
		return float64(nums[mid]+nums[mid-1]) / 2.0
	}
}
