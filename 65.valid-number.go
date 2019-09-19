import (
	"regexp"
	"strings"
)

func isNumber(s string) bool {
	r := regexp.MustCompile(`^[+-]?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(e[+-]?[0-9]+)?$`)
	return r.MatchString(strings.TrimSpace(s))
}
