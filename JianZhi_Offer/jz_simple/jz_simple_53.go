package jz_simple

// 在排序数组中查找数组I
func search(nums []int, target int) int {
	count := 0
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			count = 1
			for i := mid - 1; i >= 0; i-- {
				if nums[i] != target {
					break
				}
				count++
			}
			for i := mid + 1; i < len(nums); i++ {
				if nums[i] != target {
					break
				}
				count++
			}
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return count
}

// II 0～n-1中缺失的数字
//
//	0,1,3,4 输出 2
func missingNumber(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j { // 当等于了，再执行一遍，看看影响不影响结果，如下，如果  nums[mid] == mid  影响了也是必须要   nums[mid] != mid 不影响
		// 所以要 加上=
		mid := (i + j) >> 1
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return i
}
