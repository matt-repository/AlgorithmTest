package Sort

//快速排序

func quickSort(arr []int) []int {
	sort(arr, 0, len(arr)-1)
	return arr
}

func sort(arr []int, low int, high int) {
	if high <= low {
		return
	}
	j := partition(arr, low, high)
	sort(arr, low, j-1)
	sort(arr, j+1, high)
}
func partition(arr []int, low int, high int) int {
	i, j := low+1, high
	for true {
		for arr[i] < arr[low] {
			i++
			if i == high {
				break
			}
		}
		for arr[low] < arr[j] {
			j--
			if j == low {
				break
			}
		}
		if i >= j {
			break
		}
		exch(arr, i, j)
	}
	exch(arr, low, j)
	return j
}
func exch(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
