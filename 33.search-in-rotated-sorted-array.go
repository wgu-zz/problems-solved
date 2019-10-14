func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}
	if start == end {
		if nums[start] == target {
			return start
		}
		return -1
	}
	mid := (start + end) / 2
	if target == nums[mid] {
		return mid
	}
	if nums[mid] >= nums[0] {
		if target < nums[mid] {
			if target < nums[0] {
				return binarySearch(nums, mid+1, end, target)
			} else {
				return binarySearch(nums, start, mid-1, target)
			}
		} else {
			return binarySearch(nums, mid+1, end, target)
		}
	} else {
		if target < nums[mid] {
			return binarySearch(nums, start, mid-1, target)
		} else {
			if target <= nums[len(nums)-1] {
				return binarySearch(nums, mid+1, end, target)
			} else {
				return binarySearch(nums, start, mid-1, target)
			}
		}
	}
}
