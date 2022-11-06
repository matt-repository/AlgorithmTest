package jz_general

import "strconv"

// 把数字翻译成字符串 有多少种

//加点东西的斐波那契数列
func translateNum(num int) int {
	numList := []byte(strconv.Itoa(num))
	a, b, c := 0, 0, 1
	for i := 0; i < len(numList); i++ {
		a, b, c = b, c, 0
		if i > 1 {
			pre := (numList[i-1]-'0')*10 + (numList[i-2] - '0')
			if pre <= 25 && pre >= 10 {
				c += a + b
				continue
			}
		}
		c += b
	}
	return c
}
