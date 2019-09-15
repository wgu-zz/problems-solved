import (
	"math"
)

func numMusicPlaylists(N int, L int, K int) int {
	if K+1 > N {
		return 0
	}
	mod := int(math.Pow(10, 9) + 7)
	dp := make([]int, N+1)
	dp[0] = 1
	prev := 1
	for i := 1; i <= L; i++ {
		for j := 1; j <= N; j++ {
			cur := dp[j]
			dp[j] = (prev * (N - j + 1)) % mod
			if j > K {
				dp[j] = (dp[j] + (cur*(j-K))%mod) % mod
			}
			prev = cur
		}
		prev = 0
	}
	return dp[N]
}
