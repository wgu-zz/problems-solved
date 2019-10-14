import "math"

func divide(dividend int, divisor int) int {
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}
	dividend32, divisor32 := int32(dividend), int32(divisor)
	res := 0
	dvd, dvs := nabs(dividend32), nabs(divisor32)
	for dvs >= dvd {
		x := 1
		d := dvs
		for d >= math.MinInt32>>1 && d<<1 >= dvd {
			d <<= 1
			x <<= 1
		}
		res += x
		dvd -= d
	}
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		return -res
	}
	return res
}

func nabs(i int32) int32 {
	if i > 0 {
		return -i
	}
	return i
}
