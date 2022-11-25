package jz_general

// 礼物的最大价值

// 回溯，效率比较低，尽量别使用
func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	tempResult := &TempResult_47{
		X:   len(grid) - 1,
		Y:   len(grid[0]) - 1,
		Num: grid[0][0],
	}
	dfs_47(grid, tempResult, 0, 0)
	return tempResult.Max
}

func dfs_47(grid [][]int, tempResult *TempResult_47, x, y int) {
	if x == tempResult.X && y == tempResult.Y && tempResult.Num > tempResult.Max {
		tempResult.Max = tempResult.Num
		return
	}

	for i := 0; i < 2; i++ {
		x_ := x + dx[i]
		y_ := y + dy[i]

		if x_ <= tempResult.X && y_ <= tempResult.Y {
			tempResult.Num += grid[x_][y_]
			dfs_47(grid, tempResult, x_, y_)
			tempResult.Num -= grid[x_][y_]
		}
	}
}

type TempResult_47 struct {
	X   int //最大横坐标
	Y   int //最大纵坐标
	Num int
	Max int
}

var (
	dx = [2]int{0, 1}
	dy = [2]int{1, 0}
)

//

func maxValue1(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i-1 >= 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			if j-1 >= 0 && i == 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
			if grid[i-1][j] > grid[i][j-1] {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += grid[i][j-1]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
