package hot

import "AlgorithmTest/LeedCode/general"

// 环形链表
func hasCycle(head *general.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next // 必须不同步，不然是环形链表也相遇不了
	for fast != slow {
		if fast == nil || slow == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
