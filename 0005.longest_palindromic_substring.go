package main

func longestPalindrome(s string) string {
	if len(s) <= 0 {
		return ""
	}
	var start, end int
	for index := range s {
		i, j := findLongest(s, index, index)
		if j-i > end-start {
			start, end = i, j
		}
		i, j = findLongest(s, index, index+1)
		if j-i > end-start {
			start, end = i, j
		}
	}
	return s[start : end+1]
}

func findLongest(s string, i, j int) (int, int) {
	for ; i >= 0 && j < len(s); i, j = i-1, j+1 {
		if s[i] != s[j] {
			break
		}
	}
	return i + 1, j - 1
}
