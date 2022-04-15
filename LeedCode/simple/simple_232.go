package simple

type MyQueue struct {
	Stack []int
}

func Constructor() MyQueue {

	return MyQueue{
		Stack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.Stack = append(this.Stack, x)
}

func (this *MyQueue) Pop() int {
	result := this.Stack[0]
	this.Stack = this.Stack[1:]
	return result
}

func (this *MyQueue) Peek() int {
	return this.Stack[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.Stack) == 0 {
		return true
	} else {
		return false
	}
}

//这里为标准写法

//
//type CQueue struct {
//	Stack1 []int
//	Stack2 []int
//}
//
//
//func Constructor() CQueue {
//	return CQueue{
//		Stack1:[]int{},
//		Stack2:[]int{},
//	}
//}
//
//
//func (this *CQueue) AppendTail(value int)  {
//
//	if len(this.Stack2)!=0{
//		for len(this.Stack2)!=0{
//			this.Stack1=append(this.Stack1,this.Stack2[len(this.Stack2)-1])
//			this.Stack2= this.Stack2[:len(this.Stack2)-1]
//		}
//	}
//	this.Stack1=append(this.Stack1,value)
//
//}
//
//
//func (this *CQueue) DeleteHead() int {
//	if len(this.Stack1)!=0{
//		for len(this.Stack1)!=0{
//			this.Stack2=append(this.Stack2,this.Stack1[len(this.Stack1)-1])
//			this.Stack1= this.Stack1[:len(this.Stack1)-1]
//		}
//	}
//	if len(this.Stack2)==0{
//		return -1
//	}
//	result:=this.Stack2[len(this.Stack2)-1]
//	this.Stack2=this.Stack2[:len(this.Stack2)-1]
//	return result
//}
