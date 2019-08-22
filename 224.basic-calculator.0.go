func calculate(s string) int {
	res, sign, num := 0, 1, 0
	stack := []int{}
	signStack := []int{}
	for i, c := range s {
		if c >= '0' && c <= '9' {
			num = 10*num + int(c-'0')
		}
		if c == '+' || c == '-' {
			res += sign * num
			num = 0
			if c == '+' {
				sign = 1
			} else {
				sign = -1
			}
		}
		if c == '(' {
			stack = append(stack, res)
			res = 0
			signStack = append(signStack, sign)
			sign = 1
		}
		if c == ')' {
			res += sign * num
			num = res
			ls, lss := len(stack)-1, len(signStack)-1
			res = stack[ls]
			stack = stack[:ls]
			sign = signStack[lss]
			signStack = signStack[:lss]
		}
		if i == len(s)-1 {
			res += sign * num
		}
	}
	return res
}
