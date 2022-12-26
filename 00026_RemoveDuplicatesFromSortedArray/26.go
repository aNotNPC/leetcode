package removeduplicatesfromsortedarray

// removeDuplicates
// 时间复杂度 O(n²), 空间复杂度 O(1)
func removeDuplicates(nums []int) int {
	k := len(nums)
	for i := 1; i < k; {
		if nums[i] == nums[i-1] {
			t := nums[i]
			k -= 1
			for j := i; j < k; j++ {
				nums[j] = nums[j+1]
			}
			nums[k] = t
		} else {
			i++
		}
	}
	return k
}

func removeDuplicates1(nums []int) int {
	j := 0
	count := len(nums)
	for i := 1; i < count; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
