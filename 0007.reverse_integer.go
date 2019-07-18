package main

const (
	MaxInt = 1<<31 - 1
	MinInt = -1 << 31
)

func reverse(x int) int {
	res := 0
	for ; x != 0; x = x / 10 {
		if res > (MaxInt-x%10)/10 || res < (MinInt-x%10)/10 {
			return 0
		}
		res = res*10 + x%10
	}
	return res
}
