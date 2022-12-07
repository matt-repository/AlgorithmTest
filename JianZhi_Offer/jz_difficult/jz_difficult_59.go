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
	res := make([]int, 0, len(nums)-k+1)
	queue := make([]int, 0) //存的是index
	push := func(i int) {
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	for i := 0; i < k-1; i++ {
		push(i)
	}
	for i := k - 1; i < len(nums); i++ {
		push(i)
		for queue[0] < i-k+1 {
			queue = queue[1:]
		}
		res = append(res, nums[queue[0]])
	}
	return res
}

func dicesProbability(n int) []float64 {
	res := make([]float64, 6*n-n+1)
	var temp float64 = 1 / (6 * float64(n))

	for i := 0; i < n; i++ {
		for j := 0; j < 6; j++ {
			res[i+j] += temp
		}
	}
	return res
}
