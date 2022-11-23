package jz_difficult

// ReversePairs 数组中的逆序对
func ReversePairs(nums []int) int {
	count = 0 //记住，leed Code 的上面要初始化，不然会出错，
	mergeSort(nums, 0, len(nums)-1, make([]int, len(nums)))
	return count
}

var count = 0

func mergeSort(nums []int, left, right int, temp []int) {
	if left < right {
		//mid := right - (right-left)>>1  上面这个有问题，当 right=1 left=0 会无限循环下去会出错
		mid := left + (right-left)/2
		mergeSort(nums, left, mid, temp)
		mergeSort(nums, mid+1, right, temp)
		mergeCount(nums, left, mid, right, temp)
	}
}

func mergeCount(nums []int, left, mid, right int, temp []int) {
	i := left
	j := mid + 1
	t := 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[t] = nums[i]
			i++
			t++
		} else {
			count += mid - i + 1
			temp[t] = nums[j]
			j++
			t++
		}
	}
	for i <= mid {
		temp[t] = nums[i]
		i++
		t++
	}
	for j <= right {
		temp[t] = nums[j]
		j++
		t++
	}
	t = 0
	for left <= right {
		nums[left] = temp[t]
		left++
		t++
	}
}
