package general

///腐烂的橘子
//在给定的m x n网格grid中，每个单元格可以有以下三个值之一：
//
//值0代表空单元格；
//值1代表新鲜橘子；
//值2代表腐烂的橘子。
//每分钟，腐烂的橘子周围4 个方向上相邻 的新鲜橘子都会腐烂。
//
//返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回-1。

///广度优先 替换
func orangesRotting(grid [][]int) int {
	min := 0
	bad := [][2]int{}
	good := 0
	lengthX := len(grid)
	lengthY := len(grid[0])

	for i := 0; i < lengthX; i++ {
		for j := 0; j < lengthY; j++ {
			if grid[i][j] == 2 {
				bad = append(bad, [2]int{i, j})
			}
			if grid[i][j] == 1 {
				good++
			}
		}
	}

	for len(bad) > 0 {
		temp := [][2]int{}
		for i := 0; i < len(bad); i++ {
			x, y := bad[i][0], bad[i][1]
			for j := 0; j < 4; j++ {
				x_, y_ := x+dx[j], y+dy[j]
				if x_ >= 0 && x_ < lengthX && y_ >= 0 && y_ < lengthY && grid[x_][y_] == 1 {
					grid[x_][y_] = 2
					good--
					temp = append(temp, [2]int{x_, y_})
				}
			}
		}
		if len(temp) > 0 {
			min++
		}
		bad = temp
	}
	if good == 0 {
		return min
	} else {
		return -1
	}
}
