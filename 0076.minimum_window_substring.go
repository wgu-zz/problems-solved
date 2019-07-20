package main

func minWindow(s string, t string) string {
	m := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	count, has, hasM := len(m), 0, make(map[byte]int)
	res := ""
	for i, j := 0, 0; j < len(s); j++ {
		if need, ok := m[s[j]]; ok {
			hasM[s[j]]++
			if need == hasM[s[j]] {
				has++
			}
		} else {
			continue
		}
		for i < j {
			need, ok := m[s[i]]
			if !ok {
				i++
				continue
			}
			if hasM[s[i]] > need {
				hasM[s[i]]--
				i++
			} else {
				break
			}
		}
		if has == count {
			l := j - i + 1
			if res == "" || len(res) > l {
				res = s[i : j+1]
			}
		}
	}
	return res
}
