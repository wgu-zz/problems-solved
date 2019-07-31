package main

func permute(nums []int) [][]int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	return p(m)
}

func p(m map[int]bool) [][]int {
	res := [][]int{}
	for k := range m {
		if !m[k] {
			continue
		}
		m[k] = false
		sub := p(m)
		if len(sub) < 1 {
			sub = append(sub, []int{k})
		} else {
			for i := range sub {
				sub[i] = append(sub[i], k)
			}
		}
		res = append(res, sub...)
		m[k] = true
	}
	return res
}
