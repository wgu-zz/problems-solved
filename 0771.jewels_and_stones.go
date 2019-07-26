package main

func numJewelsInStones(J string, S string) int {
	num, m := 0, make(map[byte]bool, len(J))
	for i := 0; i < len(J); i++ {
		m[J[i]] = true
	}
	for i := 0; i < len(S); i++ {
		if m[S[i]] {
			num++
		}
	}
	return num
}
