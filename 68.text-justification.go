/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 *
 * https://leetcode.com/problems/text-justification/description/
 *
 * algorithms
 * Hard (23.93%)
 * Total Accepted:    102.3K
 * Total Submissions: 427.2K
 * Testcase Example:  '["This", "is", "an", "example", "of", "text", "justification."]\n16'
 *
 * Given an array of words and a width maxWidth, format the text such that each
 * line has exactly maxWidth characters and is fully (left and right)
 * justified.
 *
 * You should pack your words in a greedy approach; that is, pack as many words
 * as you can in each line. Pad extra spaces ' ' when necessary so that each
 * line has exactly maxWidth characters.
 *
 * Extra spaces between words should be distributed as evenly as possible. If
 * the number of spaces on a line do not divide evenly between words, the empty
 * slots on the left will be assigned more spaces than the slots on the right.
 *
 * For the last line of text, it should be left justified and no extra space is
 * inserted between words.
 *
 * Note:
 *
 *
 * A word is defined as a character sequence consisting of non-space characters
 * only.
 * Each word's length is guaranteed to be greater than 0 and not exceed
 * maxWidth.
 * The input array words contains at least one word.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * words = ["This", "is", "an", "example", "of", "text", "justification."]
 * maxWidth = 16
 * Output:
 * [
 * "This    is    an",
 * "example  of text",
 * "justification.  "
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input:
 * words = ["What","must","be","acknowledgment","shall","be"]
 * maxWidth = 16
 * Output:
 * [
 * "What   must   be",
 * "acknowledgment  ",
 * "shall be        "
 * ]
 * Explanation: Note that the last line is "shall be    " instead of "shall
 * be",
 * because the last line must be left-justified instead of fully-justified.
 * ⁠            Note that the second line is also left-justified becase it
 * contains only one word.
 *
 *
 * Example 3:
 *
 *
 * Input:
 * words =
 * ["Science","is","what","we","understand","well","enough","to","explain",
 * "to","a","computer.","Art","is","everything","else","we","do"]
 * maxWidth = 20
 * Output:
 * [
 * "Science  is  what we",
 * ⁠ "understand      well",
 * "enough to explain to",
 * "a  computer.  Art is",
 * "everything  else  we",
 * "do                  "
 * ]
 *
 *
 */

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	curLen := 0
	i, j := 0, 0
	for ; j < len(words); j++ {
		l := len(words[j])
		if curLen == 0 {
			curLen = l
			continue
		}
		if curLen+l+1 <= maxWidth {
			curLen += l + 1
			continue
		}
		num := j - i
		tmp := ""
		if num == 1 {
			tmp = words[i]
			spaecs := maxWidth - curLen
			for x := 0; x < spaecs; x++ {
				tmp += " "
			}
		} else {
			each := (maxWidth - curLen) / (num - 1)
			left := (maxWidth - curLen) % (num - 1)
			for ; i < j-1; i++ {
				tmp += words[i] + " "
				for x := 0; x < each; x++ {
					tmp += " "
				}
				if left > 0 {
					tmp += " "
					left--
				}
			}
			tmp += words[j-1]
		}
		res = append(res, tmp)
		curLen = len(words[j])
		i = j
	}
	if curLen > 0 {
		tmp := words[i]
		for i = i + 1; i < j; i++ {
			tmp += " " + words[i]
		}
		spaces := maxWidth - len(tmp)
		for x := 0; x < spaces; x++ {
			tmp += " "
		}
		res = append(res, tmp)
	}
	return res
}
