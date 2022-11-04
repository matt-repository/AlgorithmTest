package jz_general

//  二 叉搜索树 =》双向链表

func TreeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	dfs(root)
	head.Left = pre
	pre.Right = head
	return head
}

var pre, head *TreeNode

func dfs(cur *TreeNode) {
	if cur == nil {
		return
	}
	dfs(cur.Left)
	if pre == nil {
		head = cur
	} else {
		pre.Right = cur
	}
	cur.Left = pre
	pre = cur
	dfs(cur.Right)
}
