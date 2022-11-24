package jz_difficult

import (
	"container/heap"
	"sort"
)

//滑动窗口的最大值

// 1、大顶堆
var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	_a := h.IntSlice
	v := _a[len(a)-1]
	h.IntSlice = _a[:len(a)-1]
	return v
}

func MaxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

//单调队列

func MaxSlidingWindow1(nums []int, k int) []int {
	q := make([]int, 0)
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] { //等于也将其去掉
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	n := len(nums)
	ans := make([]int, 1, n-k+1) //第二个参数 len 第三个参数cap
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
