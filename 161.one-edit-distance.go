func isOneEditDistance(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls-lt > 1 || lt-ls > 1 {
		return false
	}
	for i := 0; i < ls && i < lt; i++ {
		if s[i] != t[i] {
			if ls == lt {
				return s[i+1:] == t[i+1:]
			}
			if ls > lt {
				return s[i+1:] == t[i:]
			}
			return s[i:] == t[i+1:]
		}
	}
	return ls != lt
}
