package jz_general

// 丑数
//只包含质因子为2，3，5 的数为丑数输入 n 输出第n个 丑数， 第一个为1

//源乘 2 、3、5 排序

func NthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	data := []int{1}
	dataA, dataB, dataC := 0, 0, 0
	for i := 1; i < n; i++ {
		dataA_, dataB_, dataC_ := data[dataA]*2, data[dataB]*3, data[dataC]*5
		min := isMin(dataA_, dataB_, dataC_)
		data = append(data, min)
		if min == dataA_ { // 以下都是if 因为这样会排除掉相同的数据  并且相同的只有两个，因为是2 3 5相乘的。并且相邻。
			dataA++
		}
		if min == dataB_ {
			dataB++
		}
		if min == dataC_ {
			dataC++
		}

	}
	return data[n-1]
}
func isMin(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		} else {
			return b
		}
	} else {
		if a > c {
			return c
		} else {
			return a
		}
	}

}

func min1(a, b, c int) int {
	if a <= b && a <= c { // 这里必须写等于号 ，不然是有问题的
		return a
	} else if b <= a && b <= c { // 这里必须写等于号 ，不然是有问题的
		return b
	} else {
		return c
	}
}
