package simple

// 二叉树的最大深度
//给定二叉树 [3,9,20,null,null,15,7]，返回它的最大深度 3
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := 1
	maxRight := 1
	if root.Left != nil {
		maxLeft += maxDepth(root.Left)
	}
	if root.Right != nil {
		maxRight += maxDepth(root.Right)
	}

	if maxLeft > maxRight {
		return maxLeft
	} else {
		return maxRight
	}
}
