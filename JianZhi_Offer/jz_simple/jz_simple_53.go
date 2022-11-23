package jz_simple

// 在排序数组中查找数组I
func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	mid := (start + end) >> 1
	for start <= end {
		mid = (start + end) >> 1
		if nums[mid] == target {
			result := 1
			i, j := mid-1, mid+1
			for {
				if (i < start || nums[i] != target) && (j > end || nums[j] != target) {
					break
				}
				if i >= start && nums[i] == target {
					i--
					result++
				}
				if j <= end && nums[j] == target {
					j++
					result++
				}
			}
			return result
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0
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
