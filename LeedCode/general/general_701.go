package general

//二叉搜索树中的插入操作

//这个值一定可以在叶子节点可存在

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			root.Left = insertIntoBST(root.Left, val)
		}
	}
	if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			root.Right = insertIntoBST(root.Right, val)
		}
	}
	return root
}
