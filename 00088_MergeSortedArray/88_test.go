package mergesortedarray

import "testing"

func TestMerge(t *testing.T) {
	var cases = []struct {
		name   string
		num1   []int
		m      int
		num2   []int
		n      int
		expect []int
	}{
		{"case1", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"case2", []int{1}, 1, []int{}, 0, []int{1}},
		{"case3", []int{0}, 0, []int{1}, 1, []int{1}},
		{"case4", []int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"case5", []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.num1, tt.m, tt.num2, tt.n)
			for i := 0; i < tt.m+tt.n; i++ {
				if tt.num1[i] != tt.expect[i] {
					t.Errorf("m:%d num2:%v n:%d expect:%v get:%v", tt.m, tt.num2, tt.n, tt.expect, tt.num1)
				}
			}
		})
	}
}

func TestMerge1(t *testing.T) {
	var cases = []struct {
		name   string
		num1   []int
		m      int
		num2   []int
		n      int
		expect []int
	}{
		{"case1", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"case2", []int{1}, 1, []int{}, 0, []int{1}},
		{"case3", []int{0}, 0, []int{1}, 1, []int{1}},
		{"case4", []int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"case5", []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			merge1(tt.num1, tt.m, tt.num2, tt.n)
			for i := 0; i < tt.m+tt.n; i++ {
				if tt.num1[i] != tt.expect[i] {
					t.Errorf("m:%d num2:%v n:%d expect:%v get:%v", tt.m, tt.num2, tt.n, tt.expect, tt.num1)
				}
			}
		})
	}
}
