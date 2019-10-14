import "strings"

func removeComments(source []string) []string {
	res := []string{}
	inBlock := false
	var newLine strings.Builder
	for _, s := range source {
		l := len(s)
		for i := 0; i < l; i++ {
			if inBlock && i < l-1 && s[i] == '*' && s[i+1] == '/' {
				inBlock = false
				i++
			} else if !inBlock && i < l-1 && s[i] == '/' && s[i+1] == '*' {
				inBlock = true
				i++
			} else if !inBlock && i < l-1 && s[i] == '/' && s[i+1] == '/' {
				break
			} else if !inBlock {
				newLine.WriteByte(s[i])
			}
		}
		if !inBlock && newLine.Len() > 0 {
			res = append(res, newLine.String())
			newLine.Reset()
		}
	}
	return res
}
