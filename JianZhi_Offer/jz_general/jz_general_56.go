package jz_general

// SingleNumber1 其他数字都出现了两次，就2个数字出现了一次，找出这2个数字

// 先全部异或到一个结果就是 那两个数的异或结果。然后找到这个值中 某一位为1 的，然后再遍历数组，异或这位，为0的是一个数，另外所有数异或就是另外一个数

func SingleNumber1(nums []int) []int {
	ret := 0
	a := 0
	b := 0

	for i := 0; i < len(nums); i++ {
		ret ^= nums[i]
	}
	h := 1
	for ret&h == 0 { // 这里必须用 ==0  因为只有 0和不是 0 但不一定是1 ，之前写了1 出问题
		h = h << 1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]&h == 0 { // 这里必须用 ==0  因为只有 0和不是 0 但不一定是1 ，之前写了1 出问题
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}

// 其他数字都出现了两次，就一个数字出现了一次，找出这个数字

func SingleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

// 其他数字都出现了三次，就这一个数字出现了一次，找出这个数字

// 这种可以 解决任意m 次数，然后 只有一个出现了一次，找出这个数字

func SingleNumber2(nums []int) int {
	count := make([]int, 32) //默认32位
	res := 0
	m := 3 //几个数相同
	for i := 0; i < len(nums); i++ {
		for j := 0; j < 32; j++ {
			count[j] += nums[i] & 1 // 这里 与 “1 比"   与 ”1 10  100  1000 比“ 要性能高些
			nums[i] >>= 1
		}
	}
	x := 1
	for i := 0; i < 32; i++ {
		res += (count[i] % m) * x
		x <<= 1

		//这一块比较难懂了 ,可以写成下面的位运算 ，其实就是 假如第二位 为7 则，+2的2次方
		//res <<= 1
		//res |= count[31-i] % m
	}
	return res
}
