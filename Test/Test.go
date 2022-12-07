package Test

func Test() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二 叉树 的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if isThis(root, p) && isThis(root, q) {
		leftResult := isThis(root.Left, p) && isThis(root.Left, q)
		rightResult := isThis(root.Left, p) && isThis(root.Left, q)
		if !leftResult || !rightResult {
			return root
		} else if leftResult {
			return lowestCommonAncestor(root.Left, p, q)
		} else {
			return lowestCommonAncestor(root.Right, p, q)
		}
	}
	return root
}

func isThis(root, point *TreeNode) bool {

	if root == nil {
		return false
	}

	if root == point {
		return true
	}
	return isThis(root.Left, point) || isThis(root.Right, point)
}
