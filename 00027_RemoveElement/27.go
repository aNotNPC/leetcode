package removeelement

// 2 2 3 3， 3
// 3， 3
// , 3
// 0,1,2,2,3,0,4,2, 2

/*
执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗： 2 MB , 在所有 Go 提交中击败了 57.50% 的用户
通过测试用例： 113 / 113
*/
func removeElement(nums []int, val int) int {
	count := len(nums)
	k := count
	for i := 0; i < k; {
		if nums[i] == val {
			t := nums[i]
			for j := i + 1; j < k; j++ {
				nums[j-1] = nums[j]
			}
			k--
			nums[k] = t
		} else {
			i++
		}
	}
	return k
}
