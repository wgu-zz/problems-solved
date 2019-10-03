func gameOfLife(board [][]int) {
	d := []int{-1, 0, 1}
	for i := range board {
		for j := range board[i] {
			lives := 0
			for _, d1 := range d {
				for _, d2 := range d {
					x, y := i+d1, j+d2
					if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || d1 == 0 && d2 == 0 {
						continue
					}
					if board[x][y] == 1 || board[x][y] == -1 {
						lives++
					}
				}
			}
			if board[i][j] == 0 {
				if lives == 3 {
					board[i][j] = 2
				}
			} else {
				if lives < 2 || lives > 3 {
					board[i][j] = -1
				}
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == -1 {
				board[i][j] = 0
			}
			if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}
