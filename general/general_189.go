package general

// leedCode 189题轮换数组
//给你一个数组，将数组 中的元素向右轮转K个位置，其中k 是非负数

func Rotate(nums []int, k int) {
	temp := make([]int, len(nums))

	for i, v := range nums {
		temp[(i+k)%len(nums)] = v //巧妙的确定每个值该在哪个位置
	}
	copy(nums, temp)
}

func Rotate1(nums []int, k int) {
	length := len(nums)
	temp := make([]int, 0, length)
	kTemp := k % length //多旋转的 nums的长度的次数相当于没旋转
	for i := length - kTemp; i < length; i++ {
		temp = append(temp, nums[i])
	}
	for i := 0; i < length-kTemp; i++ {
		temp = append(temp, nums[i])
	}
	copy(nums, temp)
}

//这种解法，当k 无限大，这里就废了()    仅记录自己的试探过程
func Rotate2(nums []int, k int) {
	length := len(nums)
	temp := make([]int, length) //slice 是引用传递的是引用
	for i := 0; i < k; i++ {
		temp[0] = nums[length-1]
		for j := 1; j < length; j++ {
			temp[j] = nums[j-1]
		}
		copy(nums, temp) //值拷贝
	}
}
