package simple

//leedcode 977 有序数组的平方
// 给你一个按非递减顺序排序的整数数组nums ，返回每个数字的平方 组成的新数组，要求也按非递减顺序排序
// 如 、[-4,-1,0,3,10] 输出 [0,1,9,16,100]
func SortedSquares(nums []int) []int {
	length := len(nums)
	lastNegative := -1
	//找出最后一个负数
	for i := 0; i < length && nums[i] < 0; i++ {
		lastNegative++
	}
	result := make([]int, 0, length)
	//从最后负数和第一个非负数往两边找，一个一个的添加数据
	for i, j := lastNegative, lastNegative+1; i >= 0 || j <= length-1; {
		if i == -1 {
			result = append(result, nums[j]*nums[j])
			j++
		} else if j == length {
			result = append(result, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] > nums[j]*nums[j] {
			result = append(result, nums[j]*nums[j])
			j++
		} else {
			result = append(result, nums[i]*nums[i])
			i--
		}
	}
	return result
}
