package jz_general

import "sort"

// 字符串的排列

//输入一个字符串，打印出该字符串中字符的所有排列 （不包含重复的元素）

func permutation(s string) []string {
	n := len(s)
	sList := []byte(s)
	sort.Slice(sList, func(i, j int) bool {
		return sList[i] < sList[j]
	})
	var ans []string
	pre := make([]byte, n)
	vir := make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(pre))
			return
		}
		for j, v := range vir {
			if v || (j > 0 && !vir[j-1] && sList[j] == sList[j-1]) {
				// 去掉上一层已经用过的
				continue
			}
			pre = append(pre, sList[j])
			vir[j] = true
			dfs(i + 1)
			pre = pre[:len(pre)-1]
			vir[j] = false
		}
	}
	dfs(0)
	return ans
}
