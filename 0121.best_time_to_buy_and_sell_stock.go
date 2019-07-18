package main

func maxProfit(prices []int) int {
	max, buy, sell := 0, 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[buy] {
			buy = i
		} else {
			sell = i
			profit := prices[sell] - prices[buy]
			if profit > max {
				max = profit
			}
		}
	}
	return max
}
