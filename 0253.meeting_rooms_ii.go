package main

import "sort"

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	rooms := []int{}
	for _, i := range intervals {
		room := findEarliest(rooms, i[0])
		if room == -1 {
			rooms = append(rooms, i[1])
		} else {
			rooms[room] = i[1]
		}
	}
	return len(rooms)
}

func findEarliest(rooms []int, nextStart int) int {
	earliest := -1
	for i, end := range rooms {
		if end <= nextStart && (earliest == -1 || end < earliest) {
			earliest = i
		}
	}
	return earliest
}
