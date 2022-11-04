package jz_general

//矩阵中的路径

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs_12(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs_12(board [][]byte, word string, i, j, k int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || k >= len(word) {
		return false
	}
	if board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	board[i][j] = 0

	result := dfs_12(board, word, i-1, j, k+1) || dfs_12(board, word, i+1, j, k+1) ||
		dfs_12(board, word, i, j-1, k+1) || dfs_12(board, word, i, j+1, k+1)
	board[i][j] = word[k]
	return result

}
