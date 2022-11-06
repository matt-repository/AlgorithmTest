package jz_general

import (
	"strconv"
	"strings"
)

// 把数组排成最小的数

func minNumber(nums []int) string {
	if nums == nil || len(nums) == 0 {
		return ""
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if !CheckIsMin(nums[j], nums[j+1]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	res := strings.Builder{}
	for i := 0; i < len(nums); i++ {
		res.WriteString(strconv.Itoa(nums[i]))
	}
	return res.String()
}

func CheckIsMin(num1, num2 int) bool {
	if num1 == 0 {
		return true
	}
	if num2 == 0 {
		return false
	}

	temp1 := num1
	temp2 := num2
	for temp1 != 0 {
		temp2 *= 10
		temp1 /= 10
	}

	temp1_ := num1
	temp2_ := num2
	for temp2_ != 0 {
		temp1_ *= 10
		temp2_ /= 10
	}

	if temp1_+num2 < temp2+num1 {
		return true
	} else {
		return false
	}

}
