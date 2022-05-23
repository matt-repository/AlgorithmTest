package Sort

//快速排序

func Quick(nums []int) []int {

	return _quick(nums, 0, len(nums)-1)

}

func _quick(nums []int, left, right int) []int {
	if left < right {
		index := partition(nums, left, right)
		//切片是值传递，但是 里面的数组是 指针地址，所以相当于修改的是同一片连续的内存
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
