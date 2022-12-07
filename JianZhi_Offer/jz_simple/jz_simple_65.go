package jz_simple

// 不用加减乘除 做加法

func add(a int, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	var res = a & b
	res <<= 1
	return add(res, a^b)
}
