package jz_simple

//对称二叉树

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root.Left, root.Right)

}

func check(left, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return check(left.Left, right.Right) && check(left.Right, right.Left)
}
