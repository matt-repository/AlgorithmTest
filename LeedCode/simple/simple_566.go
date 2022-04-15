package simple

//leed code 566题重塑矩阵
//在 MATLAB 中，有一个非常有用的函数 reshape ，它可以将一个m x n 矩阵重塑为另一个大小不同（r x c）的新矩阵，但保留其原始数据。
//给你一个由二维数组 mat 表示的m x n 矩阵，以及两个正整数 r 和 c ，分别表示想要的重构的矩阵的行数和列数。
//重构后的矩阵需要将原始矩阵的所有元素以相同的 行遍历顺序 填充。
//如果具有给定参数的 reshape 操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。
//
//输入：mat = [[1,2],[3,4]], r = 1, c = 4
//输出：[[1,2,3,4]]

//输入：mat = [[1,2],[3,4]], r = 2, c = 4
//输出：[[1,2],[3,4]]
func MatrixReshape(mat [][]int, r int, c int) [][]int {
	r1 := len(mat)    //行数
	c1 := len(mat[0]) //列数
	result := make([][]int, r)
	temp := make([]int, 0, c)
	tempIndex := 0
	if r*c != r1*c1 {
		result = mat
	} else {

		for _, v := range mat {
			for _, vItem := range v {
				if len(temp) < c {
					temp = append(temp, vItem)
				} else {
					result[tempIndex] = temp
					tempIndex++
					temp = make([]int, 0, c)
					temp = append(temp, vItem)
				}
			}
		}
		result[tempIndex] = temp //最后一次不会执行
	}

	return result
}
