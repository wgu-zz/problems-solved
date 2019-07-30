package main

import "sort"

func nextPermutation(nums []int) {
	i, m := len(nums)-1, make(map[int]int, 10)
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {
			min, index := nums[i], i
			for key, value := range m {
				if key > nums[i-1] && key < min {
					min = key
					index = value
				}
			}
			nums[index], nums[i-1] = nums[i-1], nums[index]
			break
		}
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = i
		}
	}
	if i == 0 {
		for x, y := 0, len(nums)-1; x < y; x, y = x+1, y-1 {
			nums[x], nums[y] = nums[y], nums[x]
		}
	} else {
		sort.Ints(nums[i:])
	}
}
