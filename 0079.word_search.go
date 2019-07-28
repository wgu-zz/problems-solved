package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if findWord(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func findWord(board [][]byte, word string, i, j int) bool {
	if len(word) < 1 {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[0] {
		return false
	}
	board[i][j] = '0'
	if findWord(board, word[1:], i-1, j) || findWord(board, word[1:], i+1, j) || findWord(board, word[1:], i, j-1) || findWord(board, word[1:], i, j+1) {
		return true
	}
	board[i][j] = word[0]
	return false
}
