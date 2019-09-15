func trap(height []int) int {
	stack := [][]int{}
	res := 0
	for i, h := range height {
		l := len(stack)
		if l < 1 || stack[l-1][1] > h {
			stack = append(stack, []int{i, h})
		} else {
			for len(stack) > 0 && h >= stack[len(stack)-1][1] {
				l := len(stack) - 1
				p := stack[l]
				stack = stack[:l]
				if l > 0 {
					res += (min(stack[l-1][1], h) - p[1]) * (i - stack[l-1][0] - 1)
				}
			}
			stack = append(stack, []int{i, h})
		}
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
