import "strconv"

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			dp[i+1] = dp[i]
		}
		if i > 0 {
			num, _ := strconv.Atoi(s[i-1 : i+1])
			if num > 9 && num < 27 {
				dp[i+1] += dp[i-1]
			}
		}
	}
	return dp[len(s)]
}
