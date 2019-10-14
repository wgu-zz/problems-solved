func sortColors(nums []int) {
	for i, j, x := 0, len(nums)-1, 0; x <= j; {
		if nums[x] == 0 {
			swap(nums, i, x)
			i++
			x++
		} else if nums[x] == 2 {
			swap(nums, j, x)
			j--
		} else {
			x++
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
