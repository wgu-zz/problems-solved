import "math"

func minWindow(s string, t string) string {
	res := ""
	l := math.MaxInt32
	m := make(map[byte]int)
	for _, c := range t {
		m[byte(c)]++
	}
	countMap := make(map[byte]int)
	count := 0
	for i, j := 0, 0; j < len(s); j++ {
		countMap[s[j]]++
		if countMap[s[j]] == m[s[j]] {
			count++
		}
		if count == len(m) {
			for count == len(m) {
				if countMap[s[i]] == m[s[i]] {
					count--
				}
				countMap[s[i]]--
				i++
			}
			if j-i+2 < l {
				res = s[i-1 : j+1]
				l = j - i + 2
			}
		}
	}
	return res
}
