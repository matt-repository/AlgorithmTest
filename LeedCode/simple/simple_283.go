package simple

//leed code 283 移动零

//给定一个数组nums ,编写一个函数将所有0移动到数组的末尾，同时保持非零元素的相对顺序

//这个方法 主要用到双指针，除常量外 并且没有开辟任何空间
func MoveZeroes(nums []int) {
	left, right, length := 0, 0, len(nums)
	for right < length {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

//自己写的思路 时间 O（n） 空间 O（1）
func MoveZeroes1(nums []int) {
	length := len(nums)
	temp := make([]int, length)
	j := 0
	for i := 0; i < length; {
		if nums[i] != 0 {
			temp[j] = nums[i]
			i++
			j++
		} else {
			i++
		}
	}
	copy(nums, temp)
}
