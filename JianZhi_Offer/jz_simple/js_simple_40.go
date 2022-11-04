package jz_simple

import "sort"

//最小的k个数

func getLeastNumber(arr []int, k int) []int {
	sort.Ints(arr)
	var result []int
	for i := 0; i < k; k++ {
		result = append(result, arr[i])
	}
	return result
}
