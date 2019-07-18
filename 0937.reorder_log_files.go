package main

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		l1 := strings.SplitN(logs[i], " ", 2)
		l2 := strings.SplitN(logs[j], " ", 2)
		isL1Digit, isL2Digit := isDigit(l1[1][0]), isDigit(l2[1][0])
		if isL1Digit {
			return false
		}
		if isL2Digit {
			return true
		}
		c := strings.Compare(l1[1], l2[1])
		if c < 0 {
			return true
		} else if c > 0 {
			return false
		}
		c = strings.Compare(l1[0], l2[0])
		if c < 0 {
			return true
		}
		return false
	})
	return logs
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}
