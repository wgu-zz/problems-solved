func calculate(s string) int {
	res := []int{}
	stack := [][]int{}
	opStack := []rune{}
	op, num := '+', 0
	for i, c := range s {
		if c >= '0' && c <= '9' {
			num = 10*num + int(c-'0')
		}
		if c == '(' {
			stack = append(stack, res)
			res = []int{}
			opStack = append(opStack, op)
			op = '+'
		}
		if c == ')' {
			cal(op, num, &res)
			num = sum(res)
			ls, lo := len(stack)-1, len(opStack)-1
			res = stack[ls]
			stack = stack[:ls]
			op = opStack[lo]
			opStack = opStack[:lo]
		}
		if c == '+' || c == '-' || c == '*' || c == '/' {
			cal(op, num, &res)
			num = 0
			op = c
		}
		if i == len(s)-1 {
			cal(op, num, &res)
		}
	}
	return sum(res)
}

func cal(op rune, num int, res *[]int) {
	if op == '+' {
		*res = append(*res, num)
	}
	if op == '-' {
		*res = append(*res, -num)
	}
	if op == '*' {
		(*res)[len(*res)-1] *= num
	}
	if op == '/' {
		(*res)[len(*res)-1] /= num
	}
}

func sum(a []int) int {
	res := 0
	for _, v := range a {
		res += v
	}
	return res
}
