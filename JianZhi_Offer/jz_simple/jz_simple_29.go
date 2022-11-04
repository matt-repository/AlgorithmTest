package jz_simple

// 顺时针打印矩阵

func spiralOrder(matrix [][]int) []int {
	xMax := len(matrix)
	if xMax == 0 {
		return []int{}
	}
	yMax := len(matrix[0])
	logo := make([][]bool, xMax)

	for i := 0; i < xMax; i++ {
		logo[i] = make([]bool, yMax)
	}
	lgo := false
	logo[0][0] = true
	result := []int{}
	result = append(result, matrix[0][0])
	i, j := 0, 0
	for {
		for j+1 < yMax && logo[i][j+1] == false {
			j++
			logo[i][j] = true
			lgo = true
			result = append(result, matrix[i][j])
		}

		for i+1 < xMax && logo[i+1][j] == false {
			i++
			logo[i][j] = true
			lgo = true
			result = append(result, matrix[i][j])
		}
		for j-1 >= 0 && logo[i][j-1] == false {
			j--
			logo[i][j] = true
			lgo = true
			result = append(result, matrix[i][j])
		}
		for i-1 >= 0 && logo[i-1][j] == false {
			i--
			logo[i][j] = true
			lgo = true
			result = append(result, matrix[i][j])
		}

		if lgo == false {
			break
		}
		lgo = false
	}

	return result
}
