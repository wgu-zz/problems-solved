func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[start] {
			if target > nums[mid] || target < nums[start] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[end] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}
