package jz_simple

// 合并两个排序的有序链表

///迭代
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	result := &ListNode{Val: -1, Next: &ListNode{}}
	temp := result

	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val < l2.Val) {
			temp.Next = l1
			l1.Next = l1
		} else {
			temp.Next = l2
			l2.Next = l2
		}
		temp = temp.Next
	}
	return result.Next
}

// 递归
func mergeTwoLists1(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	result := &ListNode{}
	if l1 == nil && (l2 != nil && l2.Val < l1.Val) {
		result = &ListNode{Val: l2.Val}
		result.Next = mergeTwoLists1(l1, l2.Next)
	} else {
		result = &ListNode{Val: l1.Val}
		result.Next = mergeTwoLists1(l1.Next, l2)
	}
	return result
}
