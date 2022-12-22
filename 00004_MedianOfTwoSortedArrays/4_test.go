package medianoftwosortedarrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	var cases = []struct {
		name   string
		nums1  []int
		nums2  []int
		expect float64
	}{
		{"case1", []int{1, 3}, []int{2}, 2},
		{"case2", []int{1, 2}, []int{3, 4}, 2.5},
		{"case3", []int{1, 7}, []int{2, 5}, 3.5},
		{"case4", []int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8}, 5},
		{"case5", []int{2, 4, 6, 8}, []int{1, 3, 5, 7, 9}, 5},
		{"case6", []int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9}, 5.5},
		{"case7", []int{}, []int{1}, 1},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ret := findMedianSortedArrays(tt.nums1, tt.nums2)
			if ret != tt.expect {
				t.Errorf("nums1:%v nums2:%v expect:%f get:%f", tt.nums1, tt.nums2, tt.expect, ret)
			}
		})
	}
}
