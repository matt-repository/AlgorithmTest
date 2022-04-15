package jz_simple

//在排序数组中查找数组I
func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	mid := (start + end) >> 1
	for start <= end {
		mid = (start + end) >> 1
		if nums[mid] == target {
			result := 1
			i, j := mid-1, mid+1
			for {
				if (i < 0 || nums[i] != target) && (j >= len(nums) || nums[j] != target) {
					break
				}
				if i >= 0 && nums[i] == target {
					i--
					result++
				}
				if j < len(nums) && nums[j] == target {
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

//
// II 0～n-1中缺失的数字

func missingNumber(nums []int) int {

	i, j := 0, len(nums)-1

	for i <= j {
		mid := (i + j) >> 1
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return i
}
