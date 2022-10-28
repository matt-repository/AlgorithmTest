package jz_general

// 数值的整数次方
// 即计算x的 n 次幂 函数 ，不得使用库函数，同时不需要考虑 大数问题

func myPow(x float64, n int) float64 {

	if x == 0 {
		return 0
	}
	var result float64 = 1
	if n < 0 {
		x = 1 / x
		n = -n
	}

	for n != 0 {
		b := n & 1
		if b == 1 {
			result *= x // 计算完了，赋值到 result上
		}
		x *= x
		n = n >> 1
	}
	return result
}
