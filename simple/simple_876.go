package simple

//leed code 876题，链表的中间结点

//给定一个头结点为 head 的非空单链表，返回链表的中间结点。
//如果有两个中间结点，则返回第二个中间结点。

//Definition for singly-linked list.

func MiddleNode(head *ListNode) *ListNode {
	result := head
	current := head
	for current != nil && current.Next != nil {
		result = result.Next
		current = current.Next.Next
	}
	return result
}

func MiddleNode1(head *ListNode) *ListNode {
	count := 1
	current := head
	for current.Next != nil {
		count++
		current = current.Next
	}
	current = head
	for i := 0; i < count/2+1; i++ {
		current = current.Next
	}
	return current
}
