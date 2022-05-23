package jz_simple

//环形链表解法,一次一次的进去，会浪费很多时间
func lastRemaining(n int, m int) int {
	data := &TempList{Value: 0}

	temp := data
	for i := 1; i < n; i++ {
		temp.Next = &TempList{
			Value: i,
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
	return temp.Value
}

type TempList struct {
	Value int
	Next  *TempList
}
