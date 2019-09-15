func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	insertNew(intervals, &res, 0, newInterval)
	return res
}

func insertNew(intervals [][]int, res *[][]int, i int, newInterval []int) {
	if i == len(intervals) {
		*res = append(*res, newInterval)
		return
	}
	if hasOverlap(intervals[i], newInterval) {
		newInterval = merge(intervals[i], newInterval)
		insertNew(intervals, res, i+1, newInterval)
	} else {
		if intervals[i][0] < newInterval[0] {
			*res = append(*res, intervals[i])
			insertNew(intervals, res, i+1, newInterval)
		} else {
			*res = append(*res, newInterval)
			for j := i; j < len(intervals); j++ {
				*res = append(*res, intervals[j])
			}
		}
	}
}

func hasOverlap(i1, i2 []int) bool {
	if i1[0] > i2[0] {
		return hasOverlap(i2, i1)
	}
	if i1[1] >= i2[0] {
		return true
	}
	return false
}

func merge(i1, i2 []int) []int {
	i1[0] = min(i1[0], i2[0])
	i1[1] = max(i1[1], i2[1])
	return i1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
