package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	var cases = []struct {
		name   string
		nums   []int
		target int
		expect int
	}{
		{"case1", []int{1, 3, 5, 6}, 5, 2},
		{"case2", []int{1, 3, 5, 6}, 2, 1},
		{"case3", []int{1, 3, 5, 6}, 7, 4},
		{"case4", []int{}, 7, 0},
		{"case5", []int{2}, 7, 1},
		{"case6", []int{2}, 2, 0},
		{"case7", []int{2}, 1, 0},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result := searchInsert(tt.nums, tt.target)
			if result != tt.expect {
				t.Errorf("nums:%v target:%d expect:%d but get:%d", tt.nums, tt.target, tt.expect, result)
			}
		})
	}
}

func TestSearchInsert1(t *testing.T) {
	var cases = []struct {
		name   string
		nums   []int
		target int
		expect int
	}{
		{"case1", []int{1, 3, 5, 6}, 5, 2},
		{"case2", []int{1, 3, 5, 6}, 2, 1},
		{"case3", []int{1, 3, 5, 6}, 7, 4},
		{"case4", []int{}, 7, 0},
		{"case5", []int{2}, 7, 1},
		{"case6", []int{2}, 2, 0},
		{"case7", []int{2}, 1, 0},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result := searchInsert1(tt.nums, tt.target)
			if result != tt.expect {
				t.Errorf("nums:%v target:%d expect:%d but get:%d", tt.nums, tt.target, tt.expect, result)
			}
		})
	}
}
