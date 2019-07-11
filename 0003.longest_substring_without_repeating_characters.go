package main

func lengthOfLongestSubstring(s string) int {
	res, head, m := 0, 0, make(map[rune]int)
	for i, v := range s {
		index, ok := m[v]
		if ok && index >= head {
			head = index + 1
		}
		m[v] = i
		if i-head+1 > res {
			res = i - head + 1
		}
	}
	return res
}
