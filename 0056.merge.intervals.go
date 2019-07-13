package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	for i, interval := range intervals {
		if i == 0 {
			res = append(res, interval)
			continue
		}
		index := len(res) - 1
		if interval[0] > res[index][1] {
			res = append(res, interval)
			continue
		}
		if interval[1] > res[index][1] {
			res[index][1] = interval[1]
		}
	}
	return res
}
