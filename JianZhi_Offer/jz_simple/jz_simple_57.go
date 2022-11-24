package jz_simple

//和为s的两个数字

// 双指针解决
func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right++
		}
	}
	return []int{}
}
