package main

func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	res[0] = 1
	for i := 1; i < l; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	p := 1
	for i := l - 1; i >= 0; i-- {
		res[i] *= p
		p *= nums[i]
	}
	return res
}
