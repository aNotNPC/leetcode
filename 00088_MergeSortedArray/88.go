package mergesortedarray

//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
/*
执行用时： 4 ms , 在所有 Go 提交中击败了 28.04% 的用户
内存消耗： 2.2 MB , 在所有 Go 提交中击败了 22.50% 的用户
通过测试用例： 59 / 59
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}

	numTmp := make([]int, m)
	for i := 0; i < m; i++ {
		numTmp[i] = nums1[i]
	}

	i, j := 0, 0
	for i < m && j < n {
		if numTmp[i] <= nums2[j] {
			nums1[i+j] = numTmp[i]
			i++
		} else {
			nums1[i+j] = nums2[j]
			j++
		}
	}

	if i >= m {
		for ; j < n; j++ {
			nums1[i+j] = nums2[j]
		}
	}
	if j >= n {
		for ; i < m; i++ {
			nums1[i+j] = numTmp[i]
		}
	}
}

// [0], 0, [1], 1
// [1], 1, [], 0
// [2, 4, 0, 0], 2, [1, 3], 2
/*
执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗： 2.2 MB , 在所有 Go 提交中击败了 34.76% 的用户
通过测试用例： 59 / 59
*/
func merge1(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for k := len(nums1) - 1; k >= 0; k-- {
		if i >= 0 && j >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else if j >= 0 {
			nums1[k] = nums2[j]
			j--
		}
	}
}
