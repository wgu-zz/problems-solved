func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			l := len(stack)
			if l < 1 {
				return false
			}
			lc := stack[l-1]
			if c == ')' && lc != '(' || c == ']' && lc != '[' || c == '}' && lc != '{' {
				return false
			}
			stack = stack[:l-1]
		}
	}
	return len(stack) == 0
}
