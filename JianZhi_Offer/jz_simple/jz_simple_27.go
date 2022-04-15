package jz_simple

//二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	result := &TreeNode{Val: root.Val}
	if root.Left != nil {
		result.Right = &TreeNode{Val: root.Left.Val}
	}
	if root.Right != nil {
		result.Left = &TreeNode{Val: root.Right.Val}
	}
	return result
}
