package jz_simple

// 圆圈中最后剩下的数字(一个)

// 环形链表解法,一次一次的进去，会超时

func lastRemaining(n int, m int) int {
	data := &ListNode{Val: 0}

	temp := data
	for i := 1; i < n; i++ {
		temp.Next = &ListNode{
			Val: i,
		}
		temp = temp.Next
	}
	temp.Next = data

	for temp.Next != temp {
		for i := 1; i < m; i++ {
			temp = temp.Next
		}
		temp.Next = temp.Next.Next

	}
	return temp.Val
}

//f(n,m)=(f(n-1,m)+m)%n  约瑟夫环的公式，什么原理真的看不懂
// 根据这个数学公式
//

func lastRemaining1(n int, m int) int {

	last := 0

	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
