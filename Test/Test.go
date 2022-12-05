package Test

import "fmt"

func Test() {
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
}

func reversePairs(nums []int) int {
	count = 0
	fmt.Println(mergeSort(nums, 0, len(nums)-1))

	return count
}

var count int

func mergeSort(nums []int, left, right int) []int {
	if left < right {
		mid := left + (right-left)/2

		leftResult := mergeSort(nums, left, mid)
		rightResult := mergeSort(nums, mid+1, right)

		return merge(leftResult, rightResult)
	}
	return []int{}
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

	for len(left) > 0 {
		result = append(result, left...)
	}
	for len(right) > 0 {
		result = append(result, right...)
	}

	return result
}
