func calculate(s string) int {
	num, op := 0, '+'
	stack := []int{}
	for i, c := range s {
		if c >= '0' && c <= '9' {
			num = 10*num + int(c-'0')
		}
		if c != '+' && c != '-' && c != '*' && c != '/' && i != len(s)-1 {
			continue
		}
		if op == '+' {
			stack = append(stack, num)
		}
		if op == '-' {
			stack = append(stack, -num)
		}
		if op == '*' {
			stack[len(stack)-1] *= num
		}
		if op == '/' {
			stack[len(stack)-1] /= num
		}
		op = c
		num = 0
	}
	res := 0
	for _, v := range stack {
		res += v
	}
	return res
}
