package simple

//leed code 第53道题最大子数组和
//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//子数组 是数组中的一个连续部分。

//贪心：若当前指针所指元素之前的和小于0，则丢弃当前之前的数列
//动态规划：若前一个元素大于0，则将其加到当前元素上

func MaxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//自己写的一个思路，只判断三个值
func MaxSubArray1(nums []int) int {

	///三个值的比较 一个是最大值，一个是当前正在计算一个连续组的和，一个是当前值
	current := 0
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if current < 0 {
			current = nums[i]
		} else {
			current = current + nums[i]
		}

		if max <= current && current >= nums[i] {
			max = current
		} else if current <= nums[i] && max <= nums[i] {
			max = nums[i]
		}
	}
	return max
}
