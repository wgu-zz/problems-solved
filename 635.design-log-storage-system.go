import (
	"strconv"
	"strings"
)

type LogSystem struct {
	m map[int]string
}

func Constructor() LogSystem {
	return LogSystem{make(map[int]string)}
}

func (this *LogSystem) Put(id int, timestamp string) {
	this.m[id] = timestamp
}

func (this *LogSystem) Retrieve(s string, e string, gra string) []int {
	res := []int{}
	graIndex := getGraIndex(gra)
	start, end := s[:graIndex], e[:graIndex]
	for id, t := range this.m {
		if strings.Compare(start, t[:graIndex]) != 1 && strings.Compare(end, t[:graIndex]) != -1 {
			res = append(res, id)
		}
	}
	return res
}

func getGraIndex(gra string) int {
	switch gra {
	case "Year":
		return 4
	case "Month":
		return 7
	case "Day":
		return 10
	case "Hour":
		return 13
	case "Minute":
		return 16
	case "Second":
		return 19
	default:
		return 0
	}
}

func convert(timestamp string) [6]byte {
	t := strings.Split(timestamp, ":")
	ta := [6]byte{}
	for i := 0; i < 6; i++ {
		d, _ := strconv.Atoi(t[i])
		if i == 0 {
			d -= 2000
		}
		ta[i] = byte(d)
	}
	return ta
}

/**
 * Your LogSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(id,timestamp);
 * param_2 := obj.Retrieve(s,e,gra);
 */
