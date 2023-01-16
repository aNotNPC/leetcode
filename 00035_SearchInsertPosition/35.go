package searchinsertposition

/*
执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗： 2.8 MB , 在所有 Go 提交中击败了 6.36% 的用户
通过测试用例： 64 / 64
*/
func searchInsert(nums []int, target int) int {
	i := 0
	j := len(nums)

	for {
		if i == j {
			return i
		}

		mid := (i + j) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			j = mid
		} else {
			i = mid + 1
		}
	}
}

/*
执行用时： 4 ms , 在所有 Go 提交中击败了 70.02% 的用户
内存消耗： 2.8 MB , 在所有 Go 提交中击败了 99.91% 的用户
通过测试用例： 64 / 64
*/
func searchInsert1(nums []int, target int) int {
	i := 0
	j := len(nums) - 1

	for {
		if i >= j {
			return i
		}

		mid := (i + j) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
}
