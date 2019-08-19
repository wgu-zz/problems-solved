/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 *
 * https://leetcode.com/problems/word-break/description/
 *
 * algorithms
 * Medium (35.99%)
 * Total Accepted:    362.2K
 * Total Submissions: 1M
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * Given a non-empty string s and a dictionary wordDict containing a list of
 * non-empty words, determine if s can be segmented into a space-separated
 * sequence of one or more dictionary words.
 *
 * Note:
 *
 *
 * The same word in the dictionary may be reused multiple times in the
 * segmentation.
 * You may assume the dictionary does not contain duplicate words.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "leetcode", wordDict = ["leet", "code"]
 * Output: true
 * Explanation: Return true because "leetcode" can be segmented as "leet
 * code".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "applepenapple", wordDict = ["apple", "pen"]
 * Output: true
 * Explanation: Return true because "applepenapple" can be segmented as "apple
 * pen apple".
 * Note that you are allowed to reuse a dictionary word.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * Output: false
 *
 *
 */
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	return canBreak(s, m, 0, dp)
}

func canBreak(s string, m map[string]bool, i int, dp []int) bool {
	if dp[i] != 0 {
		return dp[i] == 1
	}
	for j := i + 1; j <= len(s); j++ {
		if m[s[i:j]] && canBreak(s, m, j, dp) {
			dp[i] = 1
			return true
		}
	}
	dp[i] = -1
	return false
}
