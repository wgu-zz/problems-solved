package main

func maxSubArray(nums []int) int {
	max, cur, l := nums[0], nums[0], len(nums)
	for i := 1; i < l; i++ {
		if cur > 0 {
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		if cur > max {
			max = cur
		}
	}
	return max
}
