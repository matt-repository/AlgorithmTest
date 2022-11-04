package jz_simple

//用两个栈实现队列

type CQueue struct {
	Stack1 []int
	Stack2 []int
}

func Constructor() CQueue {
	return CQueue{
		Stack1: []int{},
		Stack2: []int{},
	}
}

// 在队列尾部插入整数
func (this *CQueue) AppendTail(value int) {

	this.Stack1 = append(this.Stack1, value)
}

// 在队列头部删除整数,若队列中没有元素，deleteHead 操作返回 -1
func (this *CQueue) DeleteHead() int {

	if len(this.Stack1) == 0 {
		return -1
	}
	for len(this.Stack1) > 0 {
		this.Stack2 = append(this.Stack2, this.Stack1[len(this.Stack1)-1])
		this.Stack1 = this.Stack1[:len(this.Stack1)-1]
	}
	result := this.Stack2[len(this.Stack2)-1]
	this.Stack2 = this.Stack2[:len(this.Stack2)-1]

	for len(this.Stack2) > 0 {
		this.Stack1 = append(this.Stack1, this.Stack2[len(this.Stack2)-1])
		this.Stack2 = this.Stack2[:len(this.Stack2)-1]
	}

	return result
}
