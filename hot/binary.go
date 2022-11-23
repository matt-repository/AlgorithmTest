package hot

//二分法  在有序数组中 查指定的值的下标 获取第一个 满足条件的下标

func Binary(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
