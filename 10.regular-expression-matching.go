func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	if lp == 0 {
		return ls == 0
	}
	fm := false
	if ls > 0 {
		if p[0] == '.' {
			fm = true
		} else {
			fm = s[0] == p[0]
		}
	}
	if lp > 1 && p[1] == '*' {
		return fm && isMatch(s[1:], p) || isMatch(s, p[2:])
	} else {
		return fm && isMatch(s[1:], p[1:])
	}
}
