type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) < 1 {
		return NumMatrix{[][]int{}}
	}
	row, column := len(matrix), len(matrix[0])
	sums := make([][]int, row)
	for i := range matrix {
		sums[i] = make([]int, column)
		for j := range matrix[i] {
			sum := matrix[i][j]
			if i > 0 {
				sum += sums[i-1][j]
			}
			if j > 0 {
				sum += sums[i][j-1]
			}
			if i > 0 && j > 0 {
				sum -= sums[i-1][j-1]
			}
			sums[i][j] = sum
		}
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := this.sums[row2][col2]
	if row1 > 0 {
		res -= this.sums[row1-1][col2]
	}
	if col1 > 0 {
		res -= this.sums[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		res += this.sums[row1-1][col1-1]
	}
	return res
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
