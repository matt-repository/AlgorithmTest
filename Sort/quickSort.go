package Sort

//快速排序

func quick(nums []int) []int {

	return _quick(nums, 0, len(nums)-1)

}

func _quick(nums []int, left, right int) []int {
	if left < right {
		index := partition(nums, left, right)
		_quick(nums, left, index-1)
		_quick(nums, index+1, right)
	}
	return nums
}

func partition(nums []int, left, right int) int {
	pivot := left
	index := left + 1

	for i := index + 1; i <= right; i++ {
		if nums[i] > nums[pivot] {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
	return index - 1
}
