package jz_simple

// 二叉搜索树的第k大节点

// 由于k 不是引用类型所以必须搞个 全局变量

func KthLargest(root *TreeNode, k int) int {
	_k = k
	res = 0
	dfs(root)
	return res
}

var res = 0
var _k = 0

func dfs(root *TreeNode) {
	if root != nil && root.Right != nil {
		dfs(root.Right)
	}
	_k--
	if _k == 0 {
		res = root.Val
	}
	if root != nil && root.Left != nil {
		dfs(root.Left)
	}
	return
}
