package simple

//leed code 206 题 反转链表
//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]

//递归
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return result
}

//只需要返回值。 （迭代）
func ReverseList1(head *ListNode) *ListNode {
	var result *ListNode
	var temp *ListNode
	for i := head; i != nil; i = i.Next {
		temp = &ListNode{Val: i.Val}
		if i != head {
			temp.Next = result
		}
		result = temp
	}
	return result
}

//暴力解决法；两次for
func ReverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	count := 1
	temp := head
	sliceCache := make([]int, 0, 1)
	sliceCache = append(sliceCache, temp.Val)
	for temp.Next != nil {
		temp = temp.Next
		count++
		sliceCache = append(sliceCache, temp.Val)
	}

	result := ListNode{Val: sliceCache[count-1], Next: &ListNode{}}
	temp = result.Next
	for i := count - 2; i >= 0; i-- {
		temp.Val = sliceCache[i]
		if i > 0 {
			temp.Next = &ListNode{}
		}
		if temp.Next != nil {
			temp = temp.Next
		}
	}
	return &result
}
