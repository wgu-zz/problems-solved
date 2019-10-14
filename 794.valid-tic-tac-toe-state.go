func validTicTacToe(board []string) bool {
	p1, p2 := 0, 0
	r, c := [3]int{}, [3]int{}
	d1, d2 := 0, 0
	winner1, winner2 := false, false
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				p1++
				r[i]++
				c[j]++
				if i == j {
					d1++
				}
				if i+j == 2 {
					d2++
				}
			} else if board[i][j] == 'O' {
				p2++
				r[i]--
				c[j]--
				if i == j {
					d1--
				}
				if i+j == 2 {
					d2--
				}
			}
			if r[i] == 3 || c[j] == 3 || d1 == 3 || d2 == 3 {
				winner1 = true
			}
			if r[i] == -3 || c[j] == -3 || d1 == -3 || d2 == -3 {
				winner2 = true
			}
		}
	}
	return winner1 && !winner2 && p1 == p2+1 || winner2 && !winner1 && p1 == p2 || !winner1 && !winner2 && (p1 == p2 || p1 == p2+1)
}
