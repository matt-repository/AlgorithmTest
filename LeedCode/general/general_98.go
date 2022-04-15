package general

import "math"

//98. 验证二叉搜索树

//中序遍历 的方法实现

var pre int = -1 << 63

func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !IsValidBST(root.Left) {
		return false
	}

	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return IsValidBST(root.Right)
}

//递归实现，
func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root, math.MinInt64, math.MaxInt64)
}
func check(root *TreeNode, lower, higher int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= higher {
		return false
	}

	return check(root.Left, lower, root.Val) && check(root.Right, root.Val, higher)
}
