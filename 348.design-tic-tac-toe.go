type TicTacToe struct {
	r      []int
	c      []int
	d1, d2 int
	n      int
	winner int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	return TicTacToe{make([]int, n), make([]int, n), 0, 0, n, 0}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	if this.winner != 0 {
		return this.winner
	}
	num := getNum(player)
	this.r[row] += num
	this.c[col] += num
	if row == col {
		this.d1 += num
	}
	if row+col == this.n-1 {
		this.d2 += num
	}
	if abs(this.r[row]) == this.n || abs(this.c[col]) == this.n || abs(this.d1) == this.n || abs(this.d2) == this.n {
		this.winner = player
		return player
	}
	return 0
}

func getNum(player int) int {
	if player == 1 {
		return 1
	}
	return -1
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
