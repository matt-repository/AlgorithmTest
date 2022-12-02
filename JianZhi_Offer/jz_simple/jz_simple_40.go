package jz_simple

// 最小的k个数

// 快速排序原理解决 比 排序后去拿减少一部分性能

func getLeastNumbers(arr []int, k int) []int {
	_quick(arr, 0, len(arr)-1, k)

	return arr[:k]
}

func _quick(arr []int, left, right, k int) {
	if left < right {
		index := partion(arr, left, right)
		if index == k {
			return
		}
		_quick(arr, left, index-1, k)
		_quick(arr, index+1, right, k)
	}

}

func partion(arr []int, left, right int) int {
	poivt := left
	index := left + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[poivt] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[poivt], arr[index-1] = arr[index-1], arr[poivt]
	return index - 1
}
