package general

import "math"

//三角形最小路径和

//动态规划 解决
func minimumTotal(triangle [][]int) int {
	row := len(triangle)
	col := len(triangle[row-1])
	if row == 1 {
		return triangle[0][0]
	}
	dp := make([]int, col)
	dp[0] = triangle[0][0] + triangle[1][0]
	dp[1] = triangle[0][0] + triangle[1][1]
	for i := 2; i < row; i++ {
		length := len(triangle[i])
		dp[length-1] = triangle[i][length-1] + dp[length-2]
		for j := length - 2; j >= 1; j-- {
			dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] = triangle[i][0] + dp[0]
	}
	min := math.MaxInt
	for i := 0; i < col; i++ {
		if dp[i] < min {
			min = dp[i]
		}
	}
	return min
}
