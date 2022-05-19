package general

import "math"

//计数质数

//给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
func countPrimes(n int) int {

	return 1
}

//判断是否为素数
func isTrue(num int) bool {
	if num == 2 || num == 3 {
		return true
	}
	if num%6 != 1 && num%6 != 5 {
		return false
	}
	temp := math.Sqrt(float64(num))
	//在6的倍数两侧的也可能不是质数
	for i := 5; float64(i) <= temp; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	//排除所有，剩余的是质数
	return true
}
