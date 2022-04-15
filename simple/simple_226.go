package simple

//leedcode 226题，翻转二叉树

func InvertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}
