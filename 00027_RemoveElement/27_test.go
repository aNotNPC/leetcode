package removeelement

import "testing"

func TestRemoveElement(t *testing.T) {
	var cases = []struct {
		name       string
		nums       []int
		val        int
		expectVal  int
		expectNums []int
	}{
		{"case1", []int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{"case2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			k := removeElement(tt.nums, tt.val)
			t.Logf("%+v, get:%d", tt, k)
			if k != tt.expectVal {
				t.Errorf("%+v, get:%d", tt, k)
				return
			}
			for i := 0; i < k; i++ {
				if tt.expectNums[i] != tt.nums[i] {
					t.Errorf("%+v", tt)
					return
				}
			}
		})
	}
}

func TestRemoveElement1(t *testing.T) {
	var cases = []struct {
		name       string
		nums       []int
		val        int
		expectVal  int
		expectNums []int
	}{
		{"case1", []int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{"case2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			k := removeElement1(tt.nums, tt.val)
			t.Logf("%+v, get:%d", tt, k)
			if k != tt.expectVal {
				t.Errorf("%+v, get:%d", tt, k)
				return
			}
			for i := 0; i < k; i++ {
				if tt.expectNums[i] != tt.nums[i] {
					t.Errorf("%+v", tt)
					return
				}
			}
		})
	}
}
