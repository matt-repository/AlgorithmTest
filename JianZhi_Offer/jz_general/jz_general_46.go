package jz_general

import "strconv"

// 把数字翻译成字符串 有多少种

// 加点东西的斐波那契数列
func translateNum(num int) int {
	numList := []byte(strconv.Itoa(num))
	a, b := 1, 1
	for i := 1; i < len(numList); i++ {
		pre := (numList[i] - '0') + (numList[i-1]-'0')*10
		if pre <= 25 && pre >= 10 {
			a, b = b, a+b
		} else {
			a = b
		}
	}
	return b
}

func translateNum1(num int) int {
	numStr := strconv.Itoa(num)
	p, q := 1, 1
	for i := 0; i < len(numStr); i++ {
		value1 := numStr[i-1]
		value2 := numStr[i]
		if value1 == '1' || (value1 == '2' && value2 <= '5') {
			p, q = q, p+q
		} else {
			p = q
		}
	}
	return q
}
