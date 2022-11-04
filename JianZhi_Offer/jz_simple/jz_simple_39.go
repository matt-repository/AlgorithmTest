package jz_simple

//数组中出现次数超过一半的数字

func majorityElement(nums []int) int {
	count := 0
	num := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			count++
		} else {
			if count == 0 {
				count++
				num = nums[i]
			} else {
				count--
			}
		}
	}
	return num
}
