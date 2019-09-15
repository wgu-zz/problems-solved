func findAnagrams(s string, p string) []int {
	m := make(map[byte]int)
	for _, c := range p {
		m[byte(c)]++
	}
	res := []int{}
	countMap := make(map[byte]int)
	count := 0
	for i, j := 0, 0; j < len(s); j++ {
		countMap[s[j]]++
		if countMap[s[j]] == m[s[j]] {
			count++
		}
		if j-i+1 == len(p) {
			if count == len(m) {
				res = append(res, i)
			}
			if countMap[s[i]] == m[s[i]] {
				count--
			}
			countMap[s[i]]--
			i++
		}
	}
	return res
}
