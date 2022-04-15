package simple

//101. 对称二叉树

//递归实现 ，迭代也可以实现，这里是递归，
//递归就是点点判断，走一次就好，迭代就是放到队列里。然后再判断，
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return check(root.Left, root.Right)
}

func check(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && check(node1.Left, node2.Right) && check(node1.Right, node2.Left)

}

//自己的思路，先把两边分别排序到数组里然后进行判断、中左右，中右左，注意，null要为0添加进去
func IsSymmetric1(root *TreeNode) bool {
	if root == nil {
		return false
	}

	left := transform(root.Left, true)
	right := transform(root.Right, false)

	if len(left) != len(right) {
		return false

	}

	if (left == nil) != (right == nil) {
		return false
	}

	for i, v := range left {
		if v != right[i] {
			return false
		}
	}
	return true
}

func transform(root *TreeNode, isleft bool) []int {
	result := []int{}
	if root == nil {
		result = append(result, 0)
		return result
	}
	result = append(result, root.Val)
	if isleft {
		result = append(result, transform(root.Left, isleft)...)
		result = append(result, transform(root.Right, isleft)...)

	} else {
		result = append(result, transform(root.Right, isleft)...)
		result = append(result, transform(root.Left, isleft)...)
	}
	return result
}
