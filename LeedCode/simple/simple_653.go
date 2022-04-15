package simple

//两数之和 IV
//给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

func findTarget(root *TreeNode, k int) bool {
	mapCache := map[int]struct{}{}
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		_, ok := mapCache[k-node.Val]
		if ok {
			return true
		}
		mapCache[node.Val] = struct{}{}
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}
