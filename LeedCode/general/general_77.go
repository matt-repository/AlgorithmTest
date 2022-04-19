package general

import "fmt"

//组合
//
//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
//你可以按 任何顺序 返回答案。

func Combine(n int, k int) [][]int {
	result := [][]int{}
	if k > n {
		return result
	}
	temp := &Temp_77{
		Result: [][]int{},
		Queue:  []int{},
	}
	dfs_77(n, k, 1, temp)
	return temp.Result
}

func dfs_77(n int, k int, begin int, temp *Temp_77) {
	if len(temp.Queue) == k {
		temp_ := make([]int, len(temp.Queue))
		copy(temp_, temp.Queue)
		temp.Result = append(temp.Result, temp_)
		return
	}

	for i := begin; i <= n-(k-len(temp.Queue))+1; i++ {
		temp.Queue = append(temp.Queue, i)
		fmt.Println("递归之前:", temp.Queue)
		dfs_77(n, k, i+1, temp)
		fmt.Println("递归之后:", temp.Queue)
		temp.Queue = temp.Queue[:len(temp.Queue)-1]
	}
}

type Temp_77 struct {
	Result [][]int
	Queue  []int
}
