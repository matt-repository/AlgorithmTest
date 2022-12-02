package jz_general

// 二 叉 树中和为某一值得路径

func pathSum(root *TreeNode, target int) [][]int {
	tempResult := TempResult{
		Result: [][]int{},
		Queue:  []int{},
		Num:    0,
	}
	dfs_34(root, tempResult, target)
	return tempResult.Result
}

func dfs_34(first *TreeNode, tempResult TempResult, target int) {
	if first != nil {
		tempResult.Num += first.Val
		tempResult.Queue = append(tempResult.Queue, first.Val)
		defer func() {
			tempResult.Num -= first.Val
			tempResult.Queue = tempResult.Queue[:len(tempResult.Queue)-1]
		}()
		if first.Left == nil && first.Right == nil && tempResult.Num == target {
			temp := make([]int, len(tempResult.Queue))
			copy(temp, tempResult.Queue) //这里记得一定要 temp 再添加，不然一直是同一个
			tempResult.Result = append(tempResult.Result, temp)
		} else {
			dfs_34(first.Left, tempResult, target)
			dfs_34(first.Right, tempResult, target)
		}
	}
}

type TempResult struct {
	Result [][]int
	Queue  []int
	Num    int
}
