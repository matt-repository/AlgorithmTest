package jz_simple

// 输出链表的倒数第k个节点

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	t := 0
	for fast != nil {
		if t >= k {
			slow = slow.Next
		}
		fast = fast.Next
		t++
	}
	return slow
}
