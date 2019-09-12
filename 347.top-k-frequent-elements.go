import "container/heap"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	mh := &MinHeap{}
	heap.Init(mh)
	for n, f := range m {
		heap.Push(mh, []int{n, f})
		if mh.Len() > k {
			heap.Pop(mh)
		}
	}
	res := []int{}
	for i := (*mh).Len() - 1; i >= 0; i-- {
		res = append(res, (*mh)[i][0])
	}
	return res
}

type MinHeap [][]int

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MinHeap) Less(i, j int) bool {
	if mh[i][1] == mh[j][1] {
		return mh[i][0] < mh[j][0]
	}
	return mh[i][1] < mh[j][1]
}

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.([]int))
}

func (mh *MinHeap) Pop() interface{} {
	l := len(*mh) - 1
	x := (*mh)[l]
	*mh = (*mh)[:l]
	return x
}
