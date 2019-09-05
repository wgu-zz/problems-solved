import "sort"

func accountsMerge(accounts [][]string) [][]string {
	emailToId := make(map[string]int)
	idToName := make(map[int]string)
	i := 1
	parent := make([]int, 9001)
	for _, account := range accounts {
		for j := 1; j < len(account); j++ {
			if _, ok := emailToId[account[j]]; !ok {
				emailToId[account[j]] = i
				i++
			}
			union(emailToId[account[1]], emailToId[account[j]], parent)
			idToName[emailToId[account[j]]] = account[0]
		}
	}
	idToList := make(map[int][]string)
	for email, id := range emailToId {
		leadId := find(id, parent)
		idToList[leadId] = append(idToList[leadId], email)
	}
	res := [][]string{}
	for id, list := range idToList {
		row := []string{idToName[id]}
		sort.Strings(list)
		row = append(row, list...)
		res = append(res, row)
	}
	return res
}

func find(id int, parent []int) int {
	if parent[id] == 0 {
		return id
	}
	if parent[id] != id {
		parent[id] = find(parent[id], parent)
	}
	return parent[id]
}

func union(x, y int, parent []int) {
	parent[find(y, parent)] = find(x, parent)
}
