package main

func numIslands(grid [][]byte) int {
	nums := 0
	for i, line := range grid {
		for j, v := range line {
			if v == '1' {
				nums++
				findLands(grid, i, j)
			}
		}
	}
	return nums
}

func findLands(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		findLands(grid, i-1, j)
		findLands(grid, i, j+1)
		findLands(grid, i+1, j)
		findLands(grid, i, j-1)
	}
}
