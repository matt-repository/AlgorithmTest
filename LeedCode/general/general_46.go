package general

import "fmt"

//全排列

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

func Permute(nums []int) [][]int {
	temp := &Temp_46{
		Result: [][]int{},
		Queue:  nums,
	}
	dfs_46(len(nums), 0, temp)
	return temp.Result
}

func dfs_46(n, first int, temp *Temp_46) {
	if n == first {
		queue_ := make([]int, n)
		copy(queue_, temp.Queue)
		temp.Result = append(temp.Result, queue_)
	}

	for i := first; i < n; i++ {
		temp.Queue[i], temp.Queue[first] = temp.Queue[first], temp.Queue[i]
		fmt.Println("递归之前:", temp.Queue)
		dfs_46(n, first+1, temp)
		temp.Queue[first], temp.Queue[i] = temp.Queue[i], temp.Queue[first]
		fmt.Println("递归之后:", temp.Queue)
	}
}

type Temp_46 struct {
	Result [][]int
	Queue  []int
}
