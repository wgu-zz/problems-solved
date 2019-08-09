/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 *
 * https://leetcode.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (45.75%)
 * Total Accepted:    117.6K
 * Total Submissions: 257K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * Given an encoded string, return its decoded string.
 *
 * The encoding rule is: k[encoded_string], where the encoded_string inside the
 * square brackets is being repeated exactly k times. Note that k is guaranteed
 * to be a positive integer.
 *
 * You may assume that the input string is always valid; No extra white spaces,
 * square brackets are well-formed, etc.
 *
 * Furthermore, you may assume that the original data does not contain any
 * digits and that digits are only for those repeat numbers, k. For example,
 * there won't be input like 3a or 2[4].
 *
 * Examples:
 *
 *
 * s = "3[a]2[bc]", return "aaabcbc".
 * s = "3[a2[c]]", return "accaccacc".
 * s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
 *
 *
 *
 *
 */import "strconv"

func decodeString(s string) string {
	res := []byte{}
	str, sStack, nStack := []byte{}, [][]byte{}, []int{}
	for i := 0; i < len(s); i++ {
		if isLetter(s[i]) {
			str = append(str, s[i])
		} else if isDigit(s[i]) {
			start := i
			for isDigit(s[i]) {
				i++
			}
			n, _ := strconv.Atoi(string(s[start:i]))
			nStack = append(nStack, n)
			sStack = append(sStack, str)
			str = []byte{}
		} else if s[i] == ']' {
			l := len(nStack) - 1
			num := nStack[l]
			nStack = nStack[:l]
			tmp := str
			for j := 1; j < num; j++ {
				tmp = append(tmp, str...)
			}
			l = len(sStack) - 1
			prev := sStack[l]
			sStack = sStack[:l]
			str = append(prev, tmp...)
		}
		if len(nStack) < 1 {
			res = append(res, str...)
			str = []byte{}
		}
	}
	return string(res)
}

func isLetter(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
