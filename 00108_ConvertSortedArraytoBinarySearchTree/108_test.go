package convertsortedarraytobinarysearchtree

import "testing"

func TestSortedArrayToBST(t *testing.T) {
	var cases = []struct {
		name   string
		input  []int
		except []interface{}
	}{
		{"case1", []int{-10, -3, 0, 5, 9}, []interface{}{0, -3, 9, -10, "null", 5}},
		{"case2", []int{1, 3}, []interface{}{3, 1}},
		{"case3", []int{3, 5, 8}, []interface{}{5, 3, 8}}, //[5,8,0]
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ret := SortedArrayToBST(tt.input)
			t.Logf("input:%v get:%v", tt.input, *ret)
		})
	}
}
