package jz_general

// 从上到下打印二叉树
// Ⅰ 返回一维数组
func levelOrder(root *TreeNode) []int {
	var result []int
	temp := []*TreeNode{root}
	for len(temp) > 0 {
		var tempL []*TreeNode
		for i := 0; i < len(temp); i++ {
			result = append(result, temp[i].Val)
			if temp[i].Left != nil {
				tempL = append(tempL, temp[i].Left)
			}
			if temp[i].Right != nil {
				tempL = append(tempL, temp[i].Right)
			}
		}
		temp = tempL
	}
	return result
}

// Ⅱ 返回二维数组
func levelOrder1(root *TreeNode) [][]int {
	result := [][]int{}
	temp := []*TreeNode{root}
	for len(temp) > 0 {
		var tempL []*TreeNode
		var data []int
		for i := 0; i < len(temp); i++ {
			data = append(data, temp[i].Val)
			if temp[i].Left != nil {
				tempL = append(tempL, temp[i].Left)
			}
			if temp[i].Right != nil {
				tempL = append(tempL, temp[i].Right)
			}
		}
		result = append(result, data)
		temp = tempL
	}
	return result
}

// Ⅲ 一会正序 一会倒序的返回数组

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	var temp []*TreeNode
	var logo bool
	temp = append(temp, root)

	for len(temp) > 0 {
		var tempL []*TreeNode
		var data []int
		if !logo {
			for i := 0; i < len(temp); i++ {
				data = append(data, temp[i].Val)
			}
		} else {
			for i := len(temp) - 1; i >= 0; i++ {
				data = append(data, temp[i].Val)
			}
		}
		for i := 0; i < len(temp); i++ {
			if temp[i].Left != nil {
				tempL = append(tempL, temp[i].Left)
			}
			if temp[i].Right != nil {
				tempL = append(tempL, temp[i].Right)
			}
		}
		result = append(result, data)
		logo = !logo
		temp = tempL
	}
	return result
}
