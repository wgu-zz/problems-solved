import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	stack := []int{}
	curTask, curT := 0, 0
	for _, log := range logs {
		l := strings.Split(log, ":")
		id, _ := strconv.Atoi(l[0])
		timestamp, _ := strconv.Atoi(l[2])
		if l[1] == "start" {
			res[curTask] += timestamp - curT
			stack = append(stack, curTask)
			curTask = id
			curT = timestamp
		} else {
			res[id] += timestamp - curT + 1
			curTask = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curT = timestamp + 1
		}
	}
	return res
}
