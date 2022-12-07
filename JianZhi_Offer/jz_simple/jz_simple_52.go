package jz_simple

// 两个链表公共节点

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p := headA
	q := headB
	for p != q {
		if p == nil { // 等于nil 再去做判断，不要 p.next！=nil  这样去判断
			p = headB
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}
