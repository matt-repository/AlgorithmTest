package jz_simple

// 二叉树的最近公共祖先 不存在重复值

// 二叉 搜索树 的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	result := root
	for {
		if result.Val < p.Val && result.Val < q.Val {
			result = result.Right
		} else if result.Val > p.Val && result.Val > q.Val {
			result = result.Left
		} else {
			break
		}
	}
	return result
}

// 二叉树的 最近公共祖先  不存在重复值

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left == nil {
		return right
	} else {
		return left
	}
}
