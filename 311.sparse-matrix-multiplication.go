func multiply(A [][]int, B [][]int) [][]int {
	sparseA, sparseB := getSparseArray(A), getSparseArray(B)
	row, column := len(A), len(B[0])
	res := make([][]int, row)
	for i := range res {
		res[i] = make([]int, column)
	}
	for _, a := range sparseA {
		for _, b := range sparseB {
			if a[1] == b[0] {
				res[a[0]][b[1]] += a[2] * b[2]
			}
		}
	}
	return res
}

func getSparseArray(a [][]int) [][]int {
	sparse := [][]int{}
	for i, row := range a {
		for j, n := range row {
			if n != 0 {
				sparse = append(sparse, []int{i, j, n})
			}
		}
	}
	return sparse
}
