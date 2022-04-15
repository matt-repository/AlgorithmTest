package simple

//leed code 第35道题 搜索插入位置
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
func SearchInsert(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1
	mid := (right - left) >> 1

	for left <= right {
		mid = (right + left) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { //右边
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
