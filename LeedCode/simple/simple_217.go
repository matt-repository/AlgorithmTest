package simple

import "sort"

//Leed Code 217题，
//给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
//输入：nums = [1,2,3,1]
//输出：true

func ContainsDuplicate(nums []int) bool {
	set := map[int]byte{}
	for _, v := range nums {
		if _, has := set[v]; has {
			return true
		}
		set[v] = ' '
	}
	return false
}

func ContainsDuplicate1(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
func ContainsDuplicate2(nums []int) bool {
	length := len(nums)
	for i := 0; i < length-1; i++ {

		for j := i + 1; j < length-1; j++ {
			if nums[i] == nums[j] {
				return true
			}

		}
	}
	return false
}
