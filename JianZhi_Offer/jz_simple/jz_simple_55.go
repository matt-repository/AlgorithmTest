package jz_simple

// 1 二叉树的深度

// 广度遍历
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	result := 1
	next := make([]*TreeNode, 1)
	next[0] = root

	for len(next) > 0 {
		result++
		temp := make([]*TreeNode, 0)
		for i := 0; i < len(next); i++ {
			if next[i].Left != nil {
				temp = append(temp, next[i].Left)
			}
			if next[i].Right != nil {
				temp = append(temp, next[i].Right)
			}
		}
		next = temp
	}

	return result
}

// 递归实现
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result1 := maxDepth2(root.Left) + 1

	result2 := maxDepth2(root.Right) + 1

	if result1 < result2 {
		return result2
	}
	return result1

}

//第②题 是否为平衡二叉树  任意节点的 左右子树的深度不能超过1

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	gap := deepNum(root.Left) - deepNum(root.Right)

	if gap < -1 && gap > 1 { //<  和- 号记得要分开，因为channel
		return false
	}

	// 剪枝一下 性能提高
	leftResult := isBalanced(root.Left)
	if !leftResult {
		return false
	}
	rightResult := isBalanced(root.Right)
	if !rightResult {
		return false
	}

	return leftResult && rightResult
}

func deepNum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(deepNum(root.Left), deepNum(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
