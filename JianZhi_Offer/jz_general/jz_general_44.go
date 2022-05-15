package jz_general

import "strconv"

//数字序列中某一位的数字

func FindNthDigit(n int) int {
	if n == 0 {
		return 0
	}
	start := 1
	digit := 1
	count := 9

	for n-count > 0 {
		n -= count
		start = start * 10
		digit++
		count = start * digit * 9
	}
	num := start + (n-1)/digit
	index := (n - 1) % digit
	strNum := strconv.Itoa(num)
	return int(strNum[index] - '0')
}
