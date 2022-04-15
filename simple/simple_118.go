package simple

// leedcode 118题 杨辉三脚
//给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
//在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
//输入: numRows = 5
//输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

//输入: numRows = 1
//输出: [[1]]

//官方标准
func Generate(numRows int) [][]int {

	result := make([][]int, numRows)

	for i := 1; i <= numRows; i++ {
		result[i-1] = make([]int, i)
		result[i-1][0] = 1
		result[i-1][i-1] = 1

		for j := 1; j < i-1; j++ {
			result[i-1][j] = result[i-2][j-1] + result[i-2][j]
		}
	}
	return result
}

//自己的思路写的
func Generate1(numRows int) [][]int {

	result := make([][]int, numRows)

	for i := 1; i <= numRows; i++ {
		temp := make([]int, i)
		if i == 1 {
			temp[0] = 1
		} else {
			temp[0] = 1
			for j := 1; j < i-1; j++ {
				temp[j] = result[i-2][j-1] + result[i-2][j]
			}
			temp[i-1] = 1
		}
		result[i-1] = temp
	}
	return result

}
