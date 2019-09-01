import "strings"

func isPalindrome(s string) bool {
	ls := strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		if !isLetterOrNumer(ls[i]) {
			i++
			continue
		}
		if !isLetterOrNumer(ls[j]) {
			j--
			continue
		}
		if ls[i] != ls[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isLetterOrNumer(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= '0' && b <= '9'
}
