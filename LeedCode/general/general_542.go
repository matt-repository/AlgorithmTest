package general

import "math"

//01 矩阵
//给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
//
//两个相邻元素间的距离为 1 。

//动态规划
func updateMatrix(mat [][]int) [][]int {

	maxRow := len(mat)
	maxCol := len(mat[0])

	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxCol; j++ {

			if mat[i][j] == 0 {
				continue
			}
			mat[i][j] = math.MaxInt - 10000
			if i-1 >= 0 {
				mat[i][j] = min(mat[i][j], mat[i-1][j]+1)
			}
			if j-1 >= 0 {
				mat[i][j] = min(mat[i][j], mat[i][j-1]+1)
			}
		}
	}
	for i := maxRow - 1; i >= 0; i-- {
		for j := maxCol - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				continue
			}
			if i+1 < maxRow {
				mat[i][j] = min(mat[i][j], mat[i+1][j]+1)
			}
			if j+1 < maxCol {
				mat[i][j] = min(mat[i][j], mat[i][j+1]+1)
			}

		}
	}
	return mat
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

///广度优先
func updateMatrix1(mat [][]int) [][]int {

	maxRow := len(mat)
	maxCol := len(mat[0])
	result := make([][]int, maxRow)
	for i := 0; i < maxRow; i++ {
		temp := make([]int, maxCol)
		result[i] = temp
	}
	array := [][2]int{}
	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxCol; j++ {
			result[i][j] = math.MaxInt - 10000
			if mat[i][j] == 0 {
				result[i][j] = 0
				array = append(array, [2]int{i, j})
			}
		}
	}
	for i := 0; i < len(array); i++ {
		for j := 0; j < 4; j++ {
			x, y := array[i][0]+dx[j], array[i][1]+dy[j]
			if x >= 0 && x < maxRow && y >= 0 && y < maxCol && result[x][y] > result[array[i][0]][array[i][1]]+1 {
				result[x][y] = result[array[i][0]][array[i][1]] + 1
				array = append(array, [2]int{x, y})
			}

		}
	}
	return result
}

//var (
//	dx=[4]int{0,1,0,-1}
//	dy=[4]int{1,0,-1,0}
//)
