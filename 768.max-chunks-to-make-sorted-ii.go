import "sort"

func maxChunksToSorted(arr []int) int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	sort.Ints(arrCopy)
	res := 0
	sum1, sum2 := 0, 0
	for i := range arr {
		sum1 += arr[i]
		sum2 += arrCopy[i]
		if sum1 == sum2 {
			res++
		}
	}
	return res
}
