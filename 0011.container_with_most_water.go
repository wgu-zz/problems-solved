package main

func maxArea(height []int) int {
	res := 0
	for i, j := 0, len(height)-1; i < j; {
		h := height[i]
		if height[j] < h {
			h = height[j]
		}
		area := (j - i) * h
		if area > res {
			res = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
