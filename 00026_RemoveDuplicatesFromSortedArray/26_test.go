package removeduplicatesfromsortedarray

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var cases = []struct {
		name         string
		nums         []int
		expectedNums []int
	}{
		{"case1", []int{1, 1, 2}, []int{1, 2}},
		{"case2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			k := removeDuplicates(tt.nums)
			t.Logf("nums:%v, k:%d", tt.nums, k)
			if k != len(tt.expectedNums) {
				t.Errorf("nums:%v k:%d expect:%v", tt.nums, k, tt.expectedNums)
				return
			}
			for i := 0; i < k; i++ {
				if tt.nums[i] != tt.expectedNums[i] {
					t.Errorf("nums:%v k:%d expect:%v", tt.nums, k, tt.expectedNums)
					break
				}
			}
		})
	}
}

func TestRemoveDuplicates1(t *testing.T) {
	var cases = []struct {
		name         string
		nums         []int
		expectedNums []int
	}{
		{"case1", []int{1, 1, 2}, []int{1, 2}},
		{"case2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			k := removeDuplicates1(tt.nums)
			t.Logf("nums:%v, k:%d", tt.nums, k)
			if k != len(tt.expectedNums) {
				t.Errorf("nums:%v k:%d expect:%v", tt.nums, k, tt.expectedNums)
				return
			}
			for i := 0; i < k; i++ {
				if tt.nums[i] != tt.expectedNums[i] {
					t.Errorf("nums:%v k:%d expect:%v", tt.nums, k, tt.expectedNums)
					break
				}
			}
		})
	}
}
