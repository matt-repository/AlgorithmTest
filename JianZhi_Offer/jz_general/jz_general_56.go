package jz_general

//数组中数字出现的次数
func SingleNumber(nums []int) int {
	bitSum := [32]int{}
	bitMap := 1
	for i := 0; i < len(nums); i++ {
		bitMap = 1
		for j := 0; j < 32; j++ {
			if nums[i]&bitMap != 0 {
				bitSum[j]++
			}
			bitMap <<= 1
		}
	}
	res := 0
	for i := 31; i >= 0; i-- {
		bitMap >>= 1
		if bitSum[i]%3 == 1 {
			res += bitMap
		}
	}

	return res
}
