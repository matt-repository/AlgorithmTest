package difficult

//N皇后问题

var result [][]string

func SolveNQueens(n int) [][]string {
	result = [][]string{}
	queue := make([]int, n)
	column := make(map[int]bool, n)
	left := make(map[int]bool, n)
	right := make(map[int]bool, n)

	for i := 0; i < n; i++ {
		queue[i] = -1
	}
	dp(n, 0, queue, column, left, right)
	return result
}

func dp(n, row int, queue []int, column map[int]bool, left map[int]bool, right map[int]bool) {
	if row == n {
		tempResult := getResult(queue, n)
		result = append(result, tempResult)
		return
	}

	for i := 0; i < n; i++ {
		if column[i] {
			continue
		}
		if left[i-row] {
			continue
		}
		if right[i+row] {
			continue
		}
		queue[row] = i
		column[i] = true
		left[i-row] = true
		right[i+row] = true
		dp(n, row+1, queue, column, left, right)
		queue[row] = -1
		delete(column, i)
		delete(left, i-row)
		delete(right, i+row)
	}
}

func getResult(queue []int, n int) []string {
	resultTemp := make([]string, n)
	for i := 0; i < n; i++ {
		temp := make([]byte, n)
		for j := 0; j < n; j++ {
			if queue[i] != j {
				temp[j] = '.'
			} else {
				temp[j] = 'Q'
			}
		}
		resultTemp[i] = string(temp)
	}
	return resultTemp
}
