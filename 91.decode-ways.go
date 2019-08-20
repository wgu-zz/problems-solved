import "strconv"

func numDecodings(s string) int {
	l := len(s)
	if l < 1 {
		return 0
	}
	dp := make([]int, l+1)
	dp[0] = 1
	for i := 0; i < l; i++ {
		if s[i] > '0' && s[i] <= '9' {
			dp[i+1] = dp[i]
		}
		if i < 1 {
			continue
		}
		num, _ := strconv.Atoi(s[i-1 : i+1])
		if num > 9 && num < 27 {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[l]
}
