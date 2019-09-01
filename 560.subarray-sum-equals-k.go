func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	res := 0
	sum := 0
	for _, num := range nums {
		sum += num
		res += m[sum-k]
		m[sum]++
	}
	return res
}
