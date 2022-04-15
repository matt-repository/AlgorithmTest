package simple

import "sort"

//leed code 第278道题，查出第一个错误的版本号
//输入 n=5,bad=4 输出4
//输入 n=2,bad=1 输出1
func FirstBadVersion1(n int) int {
	left := 0 //
	mid := n >> 1
	right := n
	for left < right {
		mid = (right + left) >> 1
		if !IsBadVersion(mid) { //在右边，不包括这个
			right = mid + 1
		} else { //在左边，包括这个
			left = mid
		}
	}
	// left =right =true  left-1=false
	return left
}

func FirstBadVersion2(n int) int {
	return sort.Search(n, func(n int) bool {
		return IsBadVersion(n)
	})
}

//go 自带的sort.search 源码
func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

func IsBadVersion(version int) bool {
	if version >= 3 {
		return true
	}
	return false
}
