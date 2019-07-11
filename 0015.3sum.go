package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res, l := [][]int{}, len(nums)
	for i := 0; i < l; {
		j, k := i+1, l-1
		for j < k {
			subSum := nums[j] + nums[k]
			if subSum < -nums[i] {
				j++
			} else if subSum > -nums[i] {
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
		i++
		for i < l && nums[i] == nums[i-1] {
			i++
		}
	}
	return res
}
