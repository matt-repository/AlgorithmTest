package jz_difficult

//数组中的逆序对

func ReversePairs(nums []int) int {
	mergeSort(nums)
	return count
}

var data [][2]int
var count = 0

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) >> 1
	return mergeCount(mergeSort(nums[:mid]), mergeSort(nums[mid:])) //这里是return 有返回值
	//mergeCount(mergeSort(nums[:mid]), mergeSort(nums[mid:]))  写成这样，错了，看了半天没看出来，日
	//return nums

}

func mergeCount(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	t := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result[t] = left[0]
			left = left[1:]
			t++
		} else {
			count += len(left)
			result[t] = right[0]
			right = right[1:]
			t++
		}
	}
	for len(left) > 0 {
		result[t] = left[0]
		left = left[1:]
		t++
	}
	for len(right) > 0 {
		result[t] = right[0]
		right = right[1:]
		t++
	}
	return result
}
