package general

////19. 删除链表的倒数第 N 个结点
//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//输入：head = [1], n = 1
//输出：[]
//输入：head = [1,2], n = 1
//输出：[1]

//回溯法
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	num := traverse(head, n)
	if num == n { //删除的头结点
		return head.Next
	}
	return head
}

//返回 倒数的 顺序。查到那个顺序。删除
func traverse(head *ListNode, n int) int {
	if head == nil {
		return 0
	}
	num := traverse(head.Next, n)
	if num == n {
		head.Next = head.Next.Next
	}
	return num + 1
}

//先后指针法，
func RemoveNthFromEnd1(head *ListNode, n int) *ListNode {
	first, second := head, head
	for i := 0; i < n; i++ {
		first = first.Next
	}
	if first == nil {
		return head.Next
	}
	for first != nil && first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return head
}

//找出长度，再暴力去删。最简单想到的办法
func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	count := 1
	current := head
	for current.Next != nil {
		count++
		current = current.Next
	}
	target := count - n //去除的值的下标
	temp1 := head
	if target == 0 { //当为0的时候，就去除第一个。
		return head.Next
	}
	for i := 0; i < count; i++ { //当不为0的时候，
		if i == target-1 {
			temp1.Next = temp1.Next.Next
			break
		}
		temp1 = temp1.Next
	}

	return head
}
