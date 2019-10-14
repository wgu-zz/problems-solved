import "strings"

func reverseWords(s string) string {
	res := ""
	l := len(s) - 1
	for i, j := l, l; j >= 0; {
		for i >= 0 && s[i] == ' ' {
			i--
			j--
		}
		for j >= 0 && s[j] != ' ' {
			j--
		}
		res += s[j+1:i+1] + " "
		i = j
	}
	return strings.TrimSpace(res)
}
