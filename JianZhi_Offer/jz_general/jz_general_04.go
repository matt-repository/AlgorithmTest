package jz_general

//二维数组中的查找

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		for j := len(matrix[0]) - 1; i >= 0; j++ {
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] < target {
				i--
			} else {
				j++
			}
		}
	}
	return false
}
