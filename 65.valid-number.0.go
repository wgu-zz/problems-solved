import "strings"

func isNumber(s string) bool {
	var number, dot, e, eNumber bool
	eNumber = true
	str := strings.TrimSpace(s)
	for i, c := range str {
		if c >= '0' && c <= '9' {
			number = true
			eNumber = true
		} else if c == '.' {
			if e || dot {
				return false
			}
			dot = true
		} else if c == 'e' {
			if !number || e {
				return false
			}
			eNumber = false
			e = true
		} else if c == '-' || c == '+' {
			if i != 0 && str[i-1] != 'e' {
				return false
			}
		} else {
			return false
		}
	}
	return number && eNumber
}
