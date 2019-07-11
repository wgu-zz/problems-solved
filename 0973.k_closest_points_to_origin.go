package main

func kClosest(points [][]int, K int) [][]int {
	find(points, K)
	return points[:K]
}

func find(points [][]int, k int) {
	if len(points) <= 1 {
		return
	}
	pivot := partition(points)
	leftLength := pivot + 1
	if k > leftLength {
		find(points[pivot+1:], k-leftLength)
	} else if k < leftLength {
		find(points[:pivot], k)
	}
}

func partition(points [][]int) int {
	i, j := 0, len(points)-1
	mid := (j - i) / 2
	pDist := dist(points[mid])
	swap(points, 0, mid)
	for {
		for ; i < j && dist(points[i]) <= pDist; i++ {
		}
		for ; i <= j && dist(points[j]) > pDist; j-- {
		}
		if i >= j {
			break
		}
		swap(points, i, j)
	}
	swap(points, 0, j)
	return j
}

func swap(points [][]int, i, j int) {
	tmp := points[i]
	points[i] = points[j]
	points[j] = tmp
}

func dist(i []int) int {
	return i[0]*i[0] + i[1]*i[1]
}
