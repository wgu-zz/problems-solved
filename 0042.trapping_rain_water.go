package main

func trap(height []int) int {
	var amount int
	var stack [][2]int
	for index, v := range height {
		sl := len(stack)
		if sl > 0 && v >= stack[sl-1][0] {
			flat := stack[sl-1][0]
			stack = stack[:sl-1]
			for i := sl - 2; i >= 0; i-- {
				if stack[i][0] <= v {
					amount += (stack[i][0] - flat) * (index - stack[i][1] - 1)
					flat = stack[i][0]
					stack = stack[:i]
				} else {
					amount += (v - flat) * (index - stack[i][1] - 1)
					break
				}
			}
		}
		stack = append(stack, [2]int{v, index})
	}
	return amount
}
