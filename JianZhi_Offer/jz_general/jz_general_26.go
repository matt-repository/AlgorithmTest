package jz_general

//树的子结构

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (rec(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func rec(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return rec(A.Left, B) && rec(A.Right, B)
}
