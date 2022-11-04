package jz_general

// 二 叉树搜索树的后序遍历序列
//关于二 叉搜索树，左边小于根小于右边
func verifyPostorder(postorder []int) bool {
	length := len(postorder)
	if length == 0 {
		return true
	}
	i := 0
	for ; i < length-1; i++ { //这里一定是length-1
		if postorder[i] > postorder[length-1] {
			break
		}
	}
	j := i
	for ; j < length-1; j++ { //这里一定是length-1
		if postorder[j] < postorder[length-1] {
			return false
		}
	}

	left := true
	right := true

	if i >= 0 {
		left = verifyPostorder(postorder[:i])
	}
	if i < length-1 {
		left = verifyPostorder(postorder[i : length-1])
	}
	return left && right
}
