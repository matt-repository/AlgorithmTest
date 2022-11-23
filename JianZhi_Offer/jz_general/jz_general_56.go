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
