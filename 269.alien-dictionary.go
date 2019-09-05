func alienOrder(words []string) string {
	m := make(map[byte][]byte)
	inDegree := make(map[byte]int)
	for _, word := range words {
		for _, c := range word {
			m[byte(c)] = []byte{}
			inDegree[byte(c)] = 0
		}
	}
	for i := 1; i < len(words); i++ {
		for j := 0; j < len(words[i-1]) && j < len(words[i]); j++ {
			if words[i-1][j] != words[i][j] {
				m[words[i-1][j]] = append(m[words[i-1][j]], words[i][j])
				inDegree[words[i][j]]++
				break
			}
		}
	}
	q := []byte{}
	for c, in := range inDegree {
		if in == 0 {
			q = append(q, c)
			delete(inDegree, c)
		}
	}
	res := []byte{}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		res = append(res, c)
		for _, child := range m[c] {
			inDegree[child]--
			if inDegree[child] == 0 {
				q = append(q, child)
				delete(inDegree, child)
			}
		}
	}
	if len(inDegree) > 0 {
		return ""
	}
	return string(res)
}
