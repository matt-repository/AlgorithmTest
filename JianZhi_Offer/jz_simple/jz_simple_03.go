package jz_simple

//数组中重复的数字

func findRepeatNumber(nums []int) int {
	mapCache := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		_, ok := mapCache[nums[i]]
		if ok {
			return nums[i]
		}
		mapCache[nums[i]] = true
	}
	return -1
}
