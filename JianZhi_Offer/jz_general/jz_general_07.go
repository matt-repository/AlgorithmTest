package jz_general

//重建二叉树

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	leftLength := 0
	root := &TreeNode{Val: preorder[0]}
	for i := 0; i < len(inorder); i++ {

		if preorder[0] == inorder[i] {
			leftLength = i

			break
		}
	}
	root.Left = buildTree(preorder[1:leftLength+1], inorder[:leftLength])
	root.Right = buildTree(preorder[leftLength+1:], inorder[leftLength+1:])
	return root
}
