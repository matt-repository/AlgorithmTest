package simple

//leedcode83. 删除排序链表中的重复元素
//输入：head = [1,1,2]
//输出：

//输入：head = [1,1,2,3,3]
//输出：[1,2,3]

func DeleteDuplicates(head *ListNode) *ListNode {
	for i := head; i != nil && i.Next != nil; {
		if i.Val == i.Next.Val {
			i.Next = i.Next.Next
		} else {
			i = i.Next
		}
	}

	return head
}
