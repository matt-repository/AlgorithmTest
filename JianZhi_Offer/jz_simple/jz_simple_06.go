package jz_simple

//从尾到头打印链表

//时间 O（n） 空间o（n）
func reversePrint(head *ListNode) []int {
	result := []int{}
	if head == nil {
		return result
	}
	result = append(result, reversePrint(head.Next)...)
	result = append(result, head.Val)
	return result
}

//时间 O（n） 空间o（1）
func reversePrint1(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	temp := []int{}
	for i := 0; i < len(result); i++ {
		temp = append(temp, result[i])
	}

	return temp
}
