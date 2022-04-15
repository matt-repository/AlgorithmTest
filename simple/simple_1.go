package simple

//LeedCode 第一题 两数之和
//给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

//通过hash 中不能重复的 特性解决
//时间复杂度 O(n) 空间复杂度 O(n)
func TwoSum(nums []int, target int) []int {
	mapData := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if p, ok := mapData[target-nums[i]]; ok {
			return []int{p, i}
		}
		mapData[nums[i]] = i
	}
	return []int{-1, -1}
}

//暴力解决法 时间复杂度 O(n²) 空间复杂度 O(1)
func TwoSum1(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
