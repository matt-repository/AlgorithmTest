package jz_simple

//包含min 函数的栈

type MinStack struct {
	Stack1 []int
	Stack2 []int
}

/** initialize your data structure here. */
func Constructor_MinStack() MinStack {
	return MinStack{
		Stack1: []int{},
		Stack2: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.Stack1 = append(this.Stack1, x)

	if len(this.Stack2) == 0 {
		this.Stack2 = append(this.Stack2, x)
	} else {
		min := this.Stack2[len(this.Stack2)-1]
		if min > x {
			this.Stack2 = append(this.Stack2, x)
		} else {
			this.Stack2 = append(this.Stack2, min)
		}
	}

}

func (this *MinStack) Pop() {
	this.Stack1 = this.Stack1[:len(this.Stack1)-1]
	this.Stack2 = this.Stack2[:len(this.Stack2)-1]
}

func (this *MinStack) Top() int {
	return this.Stack1[len(this.Stack1)-1]
}

func (this *MinStack) Min() int {
	return this.Stack2[len(this.Stack2)-1]
}
