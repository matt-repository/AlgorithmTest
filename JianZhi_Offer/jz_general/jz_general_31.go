package jz_general

// 栈的压入 弹出序列

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	j := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && popped[j] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}
