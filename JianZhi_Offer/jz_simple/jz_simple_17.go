package jz_simple

import (
	"strconv"
	"strings"
)

// 打印从1 到最大的n 位数

func printNumber(n int) []int {
	var result []int
	s := strings.Builder{}

	for i := 0; i < n; i++ {
		s.WriteString("9")
	}
	max, _ := strconv.Atoi(s.String())

	for i := 1; i < max+1; i++ {
		result = append(result, i)
	}
	return result
}
