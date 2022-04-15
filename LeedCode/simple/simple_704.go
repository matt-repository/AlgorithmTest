package simple

//LeetCode 第704 道题 二分查找
//给定一个n个元素有序的（升序）整型数组nums 和一个目标值target，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。
//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4

func Search1(nums []int, target int) int {
	length := len(nums)
	if length == 0 || target < nums[0] || target > nums[length-1] {
		return -1
	}
	if nums[length-1] == target {
		return length - 1
	}
	half := length / 2
	if nums[half] > target {
		leftResult := Search1(nums[0:half], target)
		return leftResult
	} else if nums[length/2] < target {
		rightResult := Search1(nums[half:length], target)
		if rightResult == -1 {
			return -1
		}
		return half + rightResult
	} else {
		return half
	}
}

func Search2(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1
	for left <= right {
		mid := (right + left) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { //在右边
			left = mid + 1
		} else { //在左边
			right = mid - 1
		}
	}

	return -1
}
