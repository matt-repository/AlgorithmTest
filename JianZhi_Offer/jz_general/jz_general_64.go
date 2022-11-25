package jz_general

// 求 1+2+。。+n
// 要求 不能使用 乘除法、for、while、if、else、switch、case 等关键字及条件判断语句 （A?B:C）

func sumNums(n int) int {
	count := 0
	var df func(i int) bool
	df = func(i int) bool {
		count += i
		return i < n && df(i+1)
	}
	df(1)
	return count
}
