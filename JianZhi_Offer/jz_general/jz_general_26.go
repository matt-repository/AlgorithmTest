package jz_general

//树的子结构

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil { // 这里是 或不是与
		return false
	}
	return rec(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func rec(A *TreeNode, B *TreeNode) bool {
	if B == nil { // 注意先后顺序，这个B!=nil 一定在前面，因为当两者都是nil 其实 是 子结构，
		// 因为这里代表下一层。不是第一层。第一层已经在上一个方法过滤掉了
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return rec(A.Left, B) && rec(A.Right, B)
}
