package jz_general

import "strconv"

//数字序列中某一位的数字
/* 数字范围    数量  位数    占多少位
   1-9        9      1       9
   10-99      90     2       180
   100-999    900    3       2700
   1000-9999  9000   4       36000  ...

   例如 2901 = 9 + 180 + 2700 + 12 即一定是4位数,第12位   n = 12;
   数据为 = 1000 + (12 - 1)/ 4  = 1000 + 2 = 1002
   定位1002中的位置 = (n - 1) %  4 = 3    s.charAt(3) = 2;
*/

func FindNthDigit(n int) int {
	if n == 0 {
		return 0
	}
	start := 1 // 数字范围开始的第一个数
	digit := 1 // n所在数字的位数
	count := 9 // 占多少位

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
