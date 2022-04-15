package general

//Leedcode 第36题，有效的数独
//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）

//一个有效的数独（部分已被填充）不一定是可解的。
//只需要根据以上规则，验证已经填入的数字是否有效即可。
//空白格用 '.' 表示。

func IsValidSudoku(board [][]byte) bool {
	rows := [9][9]int{}
	columns := [9][9]int{}
	subboxes := [3][3][9]int{}

	for i, v := range board {
		for j, jv := range v {
			if jv == '.' {
				continue
			}
			index := jv - '1' //因为上面写的空间是 9 所以下标必须为 0-8 数字为1-9所以要-1
			rows[j][index]++
			columns[i][index]++
			subboxes[i/3][j/3][index]++

			if rows[j][index] > 1 || columns[i][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

//自己按字面意思理解写的
func IsValidSudoku1(board [][]byte) bool {

	//每一行
	for _, v := range board {
		if !judge(v) {
			return false
		}
	}
	//每一列
	for i := 0; i < 9; i++ {
		temp := make([]byte, 9)
		for j := 0; j < 9; j++ {
			temp[j] = board[j][i]
		}
		if !judge(temp) {
			return false
		}
	}

	//每一方格
	tempIndex := 0
	temp := make([]byte, 9)
	for i, j := 0, 0; i < 9 && j < 9; {
		if tempIndex < 9 {
			temp[tempIndex] = board[i][j]
			tempIndex++
			if j < 9 && j != 2 && j != 5 && j != 8 {
				j++
			} else if i == 8 && (j == 2 || j == 5) {
				i = 0
				j++
			} else {
				i++
				j = j - 2
			}
		} else {
			if !judge(temp) {
				return false
			}
			tempIndex = 0
			temp = make([]byte, 9)
		}
	}
	if !judge(temp) {
		return false
	}
	return true
}

func judge(v []byte) bool {
	mapCache := map[byte]byte{}
	for index, vItem := range v {
		if vItem != '.' {
			_, ok := mapCache[vItem]
			if ok {
				return false
			}
			mapCache[vItem] = byte(index)
		}
	}
	return true
}
