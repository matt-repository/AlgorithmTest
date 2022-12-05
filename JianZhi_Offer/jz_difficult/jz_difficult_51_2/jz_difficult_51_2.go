package jz_difficult_51_2

// 逆序对
// 比较容易想到的方法 根据 归并排序

func reversePairs(nums []int) int {
	count = 0
	mergeSort(nums)
	return count
}

var count int

func mergeSort(nums []int) []int {
	if len(nums) > 1 {
		mid := len(nums) / 2
		leftResult := mergeSort(nums[:mid])
		rightResult := mergeSort(nums[mid:])
		return merge(leftResult, rightResult)
	}
	return nums
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			count += len(left)
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}
