package plusone

import "testing"

func TestPlusOne(t *testing.T) {
	var cases = []struct {
		name   string
		nums   []int
		expect []int
	}{
		{"case1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"case2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"case3", []int{0}, []int{1}},
		{"case4", []int{9}, []int{1, 0}},
		{"case5", []int{9, 9, 9, 9, 9}, []int{1, 0, 0, 0, 0, 0}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ret := plusOne(tt.nums)
			retLen := len(ret)
			expectLen := len(tt.expect)
			if retLen != expectLen {
				t.Errorf("input:%v expect:%v get:%v", tt.nums, tt.expect, ret)
				return
			}
			for i := 0; i < retLen; i++ {
				if ret[i] != tt.expect[i] {
					t.Errorf("input:%v expect:%v get:%v", tt.nums, tt.expect, ret)
					break
				}
			}
		})
	}
}
