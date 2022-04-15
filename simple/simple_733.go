package simple

//leedcode 733题 图像渲染
//有一幅以m * n的二维整数数组表示的图画image，其中image[i][j]表示该图画的像素值大小。
//
//你也被给予三个整数 sr , sc 和 newColor 。你应该从像素image[sr][sc]开始对图像进行 上色填充 。
//
//为了完成 上色工作 ，从初始像素开始，记录初始坐标的 上下左右四个方向上 像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应
//四个方向上 像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为newColor。
//
//最后返回 经过上色渲染后的图像。

//广度优先

var (
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, -1, 0, 1}
)

//广度优先
func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	lengthSr := len(image)
	lengthSc := len(image[0])
	standard := image[sr][sc]
	queue := [][]int{}
	queue = append(queue, []int{sr, sc})
	image[sr][sc] = newColor

	for i := 0; i < len(queue); i++ {
		nowX := queue[i][0]
		nowY := queue[i][1]
		for i := 0; i < 4; i++ {
			x, y := nowX+dx[i], nowY+dy[i]
			if x >= 0 && x < lengthSr && y >= 0 && y < lengthSc && image[x][y] == standard {
				image[x][y] = newColor
				queue = append(queue, []int{x, y})
			}
		}
	}
	return image
}

//深度优先
func FloodFill1(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	lengthSr := len(image)
	lengthSc := len(image[0])
	standard := image[sr][sc]
	return dfs(image, sr, sc, newColor, lengthSr, lengthSc, standard)
}

func dfs(image [][]int, sr int, sc int, newColor int, lengthSr int, lengthSc int, standard int) [][]int {
	image[sr][sc] = newColor

	for i := 0; i < 4; i++ {
		x, y := sr+dx[i], sc+dy[i]
		if x >= 0 && x < lengthSr && y >= 0 && y < lengthSc && image[x][y] == standard {
			image = dfs(image, x, y, newColor, lengthSr, lengthSc, standard)
		}
	}
	return image
}
