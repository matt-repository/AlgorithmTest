package jz_general

//构建乘积数组
//给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，
//其中B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。
//不能使用除法。

// 直观蠢办法 会很耗时

func ConstructArr(a []int) []int {
	result := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		temp := 1
		for j := 0; j < len(a); j++ {
			if i != j {
				temp *= a[j]
			}
		}
		result[i] = temp
	}
	return result
}

//时间复杂度 O(n) 空间O(1)
//当前值为 左右两边的乘积的方法

func ConstructArr1(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	res := make([]int, len(a))

	res[0] = 1
	for i := 1; i < len(a); i++ {
		res[i] = res[i-1] * a[i-1]
	}

	temp := 1
	for i := len(a) - 1; i >= 0; i-- {
		res[i] *= temp
		temp *= a[i]
	}
	return res
}
