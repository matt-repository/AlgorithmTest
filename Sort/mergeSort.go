package Sort

//归并排序

func MergeSort(nums []int) []int {
	_mergeSort(nums, 0, len(nums)-1, make([]int, len(nums)))
	return nums
}

func _mergeSort(nums []int, left, right int, temp []int) {
	if left < right {
		mid := left + (right-left)/2
		_mergeSort(nums, left, mid, temp)
		_mergeSort(nums, mid+1, right, temp)
		merge(nums, left, mid, right, temp) //合并
	}
}

func merge(nums []int, left, mid, right int, temp []int) {
	i := left
	j := mid + 1
	t := 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[t] = nums[i]
			t++
			i++
		} else {
			temp[t] = nums[j]
			t++
			j++
		}
	}
	for i <= mid {
		temp[t] = nums[i]
		t++
		i++
	}
	for j <= right {
		temp[t] = nums[j]
		t++
		j++
	}
	t = 0
	for left <= right {
		nums[left] = temp[t]
	}
}
