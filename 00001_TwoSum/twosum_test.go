package twoSum

import (
	"sort"
	"testing"
)

/*
示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/two-sum
*/
func TestTwoSum(t *testing.T) {
	var cases = []struct {
		name   string
		nums   []int
		target int
		expect []int
	}{
		{"case1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"case2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"case3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ret := twoSum(tt.nums, tt.target)
			if !isArrSame(ret, tt.expect) {
				t.Errorf("nums:%v target:%d expect:%v get:%v", tt.nums, tt.target, tt.expect, ret)
			}
		})

	}
}

func isArrSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}

func TestTwoSum1(t *testing.T) {
	var cases = []struct {
		name   string
		nums   []int
		target int
		expect []int
	}{
		{"case1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"case2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"case3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ret := twoSum1(tt.nums, tt.target)
			if !isArrSame(ret, tt.expect) {
				t.Errorf("nums:%v target:%d expect:%v get:%v", tt.nums, tt.target, tt.expect, ret)
			}
		})

	}
}

func BenchmarkTwoSum1(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9
	for i := 0; i < b.N; i++ {
		twoSum1(nums, target)
	}
}

func TestTwoSum2(t *testing.T) {
	var cases = []struct {
		name   string
		nums   []int
		target int
		expect []int
	}{
		{"case1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"case2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"case3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ret := twoSum2(tt.nums, tt.target)
			if !isArrSame(ret, tt.expect) {
				t.Errorf("nums:%v target:%d expect:%v get:%v", tt.nums, tt.target, tt.expect, ret)
			}
		})

	}
}

func BenchmarkTwoSum2(b *testing.B) {
	nums := []int{3, 3}
	target := 6
	for i := 0; i < b.N; i++ {
		twoSum2(nums, target)
	}
}
