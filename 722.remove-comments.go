import "strings"

func removeComments(source []string) []string {
	res := []string{}
	line := ""
	inBlock := false
	for _, s := range source {
		line, inBlock = pruneComment(s, line, inBlock)
		if !inBlock && len(line) > 0 {
			res = append(res, line)
			line = ""
		}
	}
	return res
}

func pruneComment(s, line string, inBlock bool) (string, bool) {
	if inBlock {
		i := strings.Index(s, "*/")
		if i == -1 {
			return line, true
		}
		return pruneComment(s[i+2:], line, false)
	}
	i, j := strings.Index(s, "/*"), strings.Index(s, "//")
	if i == -1 && j == -1 {
		return line + s, false
	}
	if i != -1 && j != -1 {
		if i < j {
			line += s[:i]
			return pruneComment(s[i+2:], line, true)
		}
		return line + s[:j], false
	}
	if i != -1 {
		line += s[:i]
		return pruneComment(s[i+2:], line, true)
	}
	return line + s[:j], false
}
