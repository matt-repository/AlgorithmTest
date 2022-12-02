package jz_general

import "sort"

// 字符串的排列

//输入一个字符串，打印出该字符串中字符的所有排列 （不包含重复的元素）

func permutation(s string) []string {
	var result []string
	sList := []byte(s)
	sort.Slice(sList, func(i, j int) bool {
		return sList[i] < sList[j]
	})
	pre := make([]byte, len(s))
	vir := make([]bool, len(s))

	var dfs func(int)
	dfs = func(i int) {
		if i == len(sList) {
			result = append(result, string(pre))
			return
		}
		for j, v := range vir {
			if v || (j > 0 && !vir[j-1] && sList[j] == sList[j-1]) {
				continue
			}
			pre[i] = sList[j]
			vir[j] = true
			dfs(i + 1)
			vir[j] = false
		}
	}
	dfs(0)
	return result
}
