package simple

//leed code 144题 二叉树的前序遍历、中序遍历、后序遍历 都是一样的道理

func PreorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	if root != nil {
		result = append(result, root.Val)
	}
	if root.Left != nil {
		result = append(result, PreorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		result = append(result, PreorderTraversal(root.Right)...)
	}

	return result
}
