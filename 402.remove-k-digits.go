func removeKdigits(num string, k int) string {
	if len(num) < 1 {
		return ""
	}
	stack := []byte{num[0]}
	for i := 1; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	for k > 0 && len(stack) > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}
	if len(stack) < 1 {
		return "0"
	}
	return string(stack)
}
