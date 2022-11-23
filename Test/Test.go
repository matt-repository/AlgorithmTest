package Test

import (
	"AlgorithmTest/JianZhi_Offer/jz_difficult"
	"fmt"
)

func Test() {

	a := jz_difficult.ReversePairs([]int{4, 5, 6, 7})
	data = make([][2]int, 0)
	mergeSort([]int{7, 5, 6, 4})
	fmt.Println(a)
	fmt.Println(count)
	fmt.Println(data)
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2

	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))

}

var data [][2]int
var count = 0

func merge(left []int, right []int) []int {
	result := make([]int, 0)

	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			for p := 0; p < len(left); p++ {
				data = append(data, [2]int{left[p], right[0]})
			}
			count += len(left)
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result

}
