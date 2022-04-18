package general

//组合
//
//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
//你可以按 任何顺序 返回答案。

func combine(n int, k int) [][]int {
	if k > n {
		return nil
	}
	result := &Result{Data: [][]int{}}
	queue := []int{}
	dfs_77(n, k, 1, queue, result)
	return result.Data
}

func dfs_77(n int, k int, begin int, queue []int, result *Result) {
	if len(queue) == k {
		temp := []int{}
		for i := 0; i < len(queue); i++ {
			temp = append(temp, queue[i])
		}
		result.Data = append(result.Data, temp)
		return
	}

	for i := begin; i <= n-(k-len(queue))+1; i++ {
		queue = append(queue, i)
		dfs_77(n, k, i+1, queue, result)
		queue = queue[:len(queue)-1]
	}
}

type Result struct {
	Data [][]int
}
