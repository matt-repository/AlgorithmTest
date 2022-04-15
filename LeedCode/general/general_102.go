package general

//leed code 102题
//二叉树的层序遍历
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	result := [][]int{}
	for queue != nil {
		temp := []int{}
		tempQueue := []*TreeNode{}
		for _, v := range queue {
			temp = append(temp, v.Val)
			if v.Left != nil {
				tempQueue = append(tempQueue, v.Left)
			}
			if v.Right != nil {
				tempQueue = append(tempQueue, v.Right)
			}

		}
		result = append(result, temp)
		queue = nil
		if len(tempQueue) > 0 {
			queue = tempQueue
		}
	}
	return result
}
