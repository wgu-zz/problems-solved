import "container/heap"

func reorganizeString(S string) string {
	m := make(map[byte]int)
	for _, c := range S {
		b := byte(c)
		m[b]++
	}
	mh := &MaxHeap{}
	heap.Init(mh)
	for b, v := range m {
		heap.Push(mh, Pair{b, v})
	}
	if (*mh)[0].Val > (len(S)+1)/2 {
		return ""
	}
	used := Pair{0, 0}
	res := make([]byte, len(S))
	for i := 0; i < len(S); i++ {
		tmp := heap.Pop(mh).(Pair)
		res[i] = tmp.B
		tmp.Val--
		if used.Val > 0 {
			heap.Push(mh, used)
		}
		used = tmp
	}
	return string(res)
}

type Pair struct {
	B   byte
	Val int
}

type MaxHeap []Pair

func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MaxHeap) Less(i, j int) bool {
	return mh[i].Val > mh[j].Val
}

func (mh *MaxHeap) Push(x interface{}) {
	*mh = append(*mh, x.(Pair))
}

func (mh *MaxHeap) Pop() interface{} {
	l := len(*mh) - 1
	x := (*mh)[l]
	*mh = (*mh)[:l]
	return x
}
