package plusone

/*
执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗： 1.9 MB , 在所有 Go 提交中击败了 8.28% 的用户
通过测试用例： 111 / 111
*/
func plusOne(digits []int) []int {
	count := len(digits)
	ret := make([]int, count)
	mod := 1
	for i := count - 1; i >= 0; i-- {
		num := digits[i] + mod
		if num == 10 {
			ret[i] = 0
			mod = 1
		} else {
			ret[i] = num
			mod = 0
		}
	}
	if mod != 0 {
		ret = append([]int{1}, ret...)
	}
	return ret
}
