package convertsortedarraytobinarysearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 输入：nums = [-10,-3,0,5,9]
// 输出：[0,-3,9,-10,null,5]
// 解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

// 执行结果： 通过 显示详情 查看示例代码 添加备注
// 执行用时： 4 ms , 在所有 Go 提交中击败了 48.03% 的用户
// 内存消耗： 3.3 MB , 在所有 Go 提交中击败了 97.08% 的用户
// 通过测试用例：31 / 31
func SortedArrayToBST(nums []int) *TreeNode {
	ret := new(TreeNode)
	count := len(nums)
	if count == 3 {
		ret.Val = nums[1]
		ret.Left = new(TreeNode)
		ret.Left.Val = nums[0]
		ret.Right = new(TreeNode)
		ret.Right.Val = nums[2]
		return ret
	} else if count == 2 {
		ret.Val = nums[1]
		ret.Left = new(TreeNode)
		ret.Left.Val = nums[0]
		return ret
	} else if count == 1 {
		ret.Val = nums[0]
		return ret
	}

	mid := (len(nums) - 1) / 2
	ret.Val = nums[mid]
	ret.Left = SortedArrayToBST(nums[:mid])
	ret.Right = SortedArrayToBST(nums[mid+1:])
	return ret
}
