package jz_generl_47

func maxValue(grid [][]int) int {
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
