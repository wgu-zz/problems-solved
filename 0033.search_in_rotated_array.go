package main

func search(nums []int, target int) int {
	l := len(nums)
	if l <= 0 {
		return -1
	}
	for i, j := 0, l-1; i <= j; {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] {
			if target > nums[mid] || target < nums[0] {
				i = mid + 1
			} else {
				j = mid - 1
			}
		} else {
			if target < nums[mid] || target >= nums[0] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
	}
	return -1
}
