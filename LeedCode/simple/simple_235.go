package simple

//235. 二叉搜索树的最近公共祖先
//给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//所有节点的值都是唯一的。
//p、q 为不同节点且均存在于给定的二叉搜索树中。

//输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
//输出: 6
//解释: 节点 2 和节点 8 的最近公共祖先是 6。

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root.Val > q.Val && root.Val > p.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < q.Val && root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root

}
