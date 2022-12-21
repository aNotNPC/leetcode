// https://leetcode.cn/problems/two-sum/

package twoSum

func twoSum(nums []int, target int) []int {
	count := len(nums)
	m := make(map[int][]int, count)
	for i, num := range nums {
		m[num] = append(m[num], i)
	}

	ret := make([]int, 0, 2)
	for k, v := range m {
		sub := target - k
		if sub == k {
			if len(v) >= 2 {
				ret = append(ret, v[0])
				ret = append(ret, v[1])
				return ret
			}
			continue
		}

		if indexs, ok := m[sub]; ok {
			ret = append(ret, v[0])
			ret = append(ret, indexs[0])
			return ret
		}
	}
	return ret
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		sub := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == sub {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, 1)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}

		m[num] = i
	}
	return nil
}
