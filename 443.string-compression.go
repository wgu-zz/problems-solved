import "strconv"

func compress(b []byte) int {
	res := 0
	if len(b) < 1 {
		return res
	}
	c := b[0]
	l := 0
	for i := 0; i <= len(b); i++ {
		if i < len(b) && b[i] == c {
			l++
		}
		if i == len(b) || b[i] != c {
			b[res] = c
			res++
			if i < len(b) {
				c = b[i]
			}
			if l > 1 {
				s := strconv.Itoa(l)
				for _, ch := range s {
					b[res] = byte(ch)
					res++
				}
				l = 1
			}
		}
	}
	return res
}
