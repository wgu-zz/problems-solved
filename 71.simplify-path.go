import "strings"

func simplifyPath(path string) string {
	stack := []string{}
	name := []byte{}
	for i, c := range path {
		if c != '/' {
			name = append(name, byte(c))
		}
		if (c == '/' || i == len(path)-1) && len(name) > 0 {
			p := string(name)
			if p == ".." {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else if p != "." {
				stack = append(stack, p)
			}
			name = []byte{}
		}
	}
	if len(stack) < 1 {
		return "/"
	}
	var sb strings.Builder
	for _, s := range stack {
		sb.WriteString("/" + s)
	}
	return sb.String()
}
