package Test

func Test() {

}

// 矩阵中的路径
func exist(board [][]byte, word string) bool {

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			if df(board, &TempStruct{[]byte{}, i, j}, word) {
				return true
			}
		}
	}

	return false
}

func df(board [][]byte, temp *TempStruct, word string) bool {
	if len(temp.Str) == len(word) && string(temp.Str) == word {
		return true
	}
	tempValue := board[temp.x][temp.y]
	board[temp.x][temp.y] = ' '
	for i := 0; i < 4; i++ {
		temp.x = temp.x + dx[i]
		temp.y = temp.y + dy[i]
		if board[temp.x][temp.y] != ' ' && temp.x >= 0 && temp.y >= 0 && temp.x < len(board) && temp.y < len(board[0]) {
			temp.Str = append(temp.Str, board[temp.x][temp.y])
			df(board, temp, word)
		}
	}
	temp.Str = temp.Str[:len(temp.Str)-1]
	board[temp.x][temp.y] = tempValue
	return false
}

var (
	dx = [4]int{0, 1, 0, -1}
	dy = [4]int{1, 0, -1, 0}
)

type TempStruct struct {
	Str []byte
	x   int
	y   int
}
