package main

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}
		last := len(stack) - 1
		if last < 0 || (s[i] == ')' && stack[last] != '(') || (s[i] == '}' && stack[last] != '{') || (s[i] == ']' && stack[last] != '[') {
			return false
		}
		stack = stack[:last]
	}
	return len(stack) == 0
}
