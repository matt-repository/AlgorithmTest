package general

//leedCode第695题，岛屿的最大面积
//岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。
//你可以假设grid 的四个边缘都被 0（代表水）包围着。
//岛屿的面积是岛上值为 1 的单元格的数目。
//计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

func MaxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			tempMax := dfs(grid, i, j)
			if tempMax > max {
				max = tempMax
			}
		}
	}
	return max
}

func dfs(grid [][]int, pointX int, pointY int) int {
	max := 0
	if pointX >= 0 && pointX < len(grid) && pointY >= 0 && pointY < len(grid[0]) && grid[pointX][pointY] == 1 {
		max++
		grid[pointX][pointY] = 0
		for j := 0; j < 4; j++ {
			x, y := dx[j]+pointX, dy[j]+pointY
			max += dfs(grid, x, y)
		}
	}
	return max
}

var (
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, -1, 0, 1}
)
