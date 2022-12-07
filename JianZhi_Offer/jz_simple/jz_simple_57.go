package jz_simple

//和为s的两个数字

// 双指针解决
func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right++
		}
	}
	return []int{}
}

// 和为s的连续正数序列,至少两个数

//输入：target=9
//输出：[[2,3,4],[4,5]]

// 滑动窗口解决

func FindContinuousSequence(target int) [][]int {
	if target < 3 {
		return [][]int{}
	}
	res := make([][]int, 0)
	left := 1
	right := 2
	count := left + right
	for left < (target+1)/2 {
		if count == target {
			temp := make([]int, right-left+1)
			j := 0
			for i := left; i <= right; i++ {
				temp[j] = i
				j++
			}
			res = append(res, temp)
			count -= left
			left++
		} else if count < target {
			right++
			count += right
		} else {
			count -= left
			left++
		}
	}
	return res
}
