/*
 * @lc app=leetcode id=224 lang=golang
 *
 * [224] Basic Calculator
 *
 * https://leetcode.com/problems/basic-calculator/description/
 *
 * algorithms
 * Hard (33.48%)
 * Total Accepted:    116.8K
 * Total Submissions: 348.3K
 * Testcase Example:  '"1 + 1"'
 *
 * Implement a basic calculator to evaluate a simple expression string.
 *
 * The expression string may contain open ( and closing parentheses ), the plus
 * + or minus sign -, non-negative integers and empty spaces  .
 *
 * Example 1:
 *
 *
 * Input: "1 + 1"
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: " 2-1 + 2 "
 * Output: 3
 *
 * Example 3:
 *
 *
 * Input: "(1+(4+5+2)-3)+(6+8)"
 * Output: 23
 * Note:
 *
 *
 * You may assume that the given expression is always valid.
 * Do not use the eval built-in library function.
 *
 *
 */
func calculate(s string) int {
	sign := 1
	stack := []int{}
	num := 0
	curSum := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = 10*num + int(c-'0')
		} else if c == '-' || c == '+' {
			curSum += sign * num
			num = 0
			if c == '-' {
				sign = -1
			} else {
				sign = 1
			}
		} else if c == '(' {
			stack = append(stack, curSum)
			curSum = 0
			stack = append(stack, sign)
			sign = 1
		} else if c == ')' {
			curSum += sign * num
			num = curSum
			l := len(stack)
			sign = stack[l-1]
			curSum = stack[l-2]
			stack = stack[:l-2]
		}
	}
	curSum += sign * num
	return curSum
}
