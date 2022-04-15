package simple

//leed code 700题 二叉树中的搜索

func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	left := searchBST(root.Left, val)
	if left != nil {
		return left
	} else {
		return searchBST(root.Right, val)
	}
}
