package general

//leed code 73题，矩阵置零
//给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
//输出：[[1,0,1],[0,0,0],[1,0,1]]
//输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
//输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

//时间复杂度 O(M*N) 空间复杂度（1）
func SetZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	row := false
	col := false
	for _, v := range matrix {
		if v[0] == 0 {
			col = true
			break
		}
	}
	for _, v := range matrix[0] {
		if v == 0 {
			row = true
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if col {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

}

//空间长度为 O(M*N) 时间复杂度 O（M*N）
func SetZeroes1(matrix [][]int) {

	target := [][]int{}
	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix[0]); j++ {

			if matrix[i][j] == 0 {
				target = append(target, []int{i, j})
			}
		}
	}

	for k := 0; k < len(target); k++ {
		i := target[k][0]
		j := target[k][1]

		for m := 0; m < len(matrix[0]); m++ {
			matrix[i][m] = 0
		}
		for n := 0; n < len(matrix); n++ {
			matrix[n][j] = 0
		}
	}

}
