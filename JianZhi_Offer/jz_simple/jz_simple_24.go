package jz_simple

//反转链表
func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	for head != nil {
		temp := &ListNode{Val: head.Val}
		if result != nil {
			temp.Next = result
		}
		result = temp
		head = head.Next
	}
	return result
}
