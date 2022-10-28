package jz_general

//二维数组中的查找

// 错误范例 （经典）
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for i := 0; i < len(matrix); {
		for j := len(matrix[0]) - 1; j >= 0; { //这里这样写不行，因为 当 i超过了 范围时。他还是会去执行，错误范例放在这
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] < target {
				i++
			} else {
				j--
			}
		}
		i++ //这里 千万别忘了 不然死循环了
	}
	return false
}

//  正确例子
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for i, j := 0, len(matrix[0])-1; i < len(matrix) && j >= 0; {
		// for 循环这里  &&  运算符 别忘了写，  初始化 j= len(matrix[0])-1  而不是  len(matrix[0])
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
