func numIslands(grid [][]byte) int {
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res++
				mark(grid, i, j)
			}
		}
	}
	return res
}

func mark(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	dir := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	for _, d := range dir {
		ii := i + d[0]
		jj := j + d[1]
		if ii >= 0 && ii < len(grid) && jj >= 0 && jj < len(grid[0]) && grid[ii][jj] == '1' {
			mark(grid, ii, jj)
		}
	}
}
