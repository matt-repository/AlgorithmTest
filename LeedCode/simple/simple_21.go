package simple

//leed code21题合并两个有序链表

//迭代写法
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}

	temp := ListNode{}
	result.Next = &temp
	for list1 != nil && list2 != nil {
		if list1 == nil {
			temp = ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		if list2 == nil {
			temp = ListNode{Val: list1.Val}
			list1 = list1.Next
		}
		if list1.Val < list2.Val {
			temp = ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			temp = ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		temp = *temp.Next
	}
	return result.Next
}

//递归写法
func MergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = MergeTwoLists1(list1.Next, list2)
		return list1

	} else {
		list2.Next = MergeTwoLists(list1, list2.Next)
		return list2
	}

}
