package main

import "sort"

type TimeMap struct {
	M map[string][]Pair
}

type Pair struct {
	Value     string
	Timestamp int
}

/** Initialize your data structure here. */
func Constructor3() TimeMap {
	return TimeMap{make(map[string][]Pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	p := Pair{value, timestamp}
	this.M[key] = append(this.M[key], p)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	i := sort.Search(len(this.M[key]), func(i int) bool {
		return timestamp < this.M[key][i].Timestamp
	})
	if i == 0 {
		return ""
	}
	return this.M[key][i-1].Value
	//return binarySearch(pairs, timestamp)
}

func binarySearch(pairs []Pair, timestamp int) string {
	l := len(pairs)
	if l < 1 {
		return ""
	}
	mid := l / 2
	if pairs[mid].Timestamp > timestamp {
		return binarySearch(pairs[:mid], timestamp)
	}
	if mid < l-1 && pairs[mid+1].Timestamp <= timestamp {
		return binarySearch(pairs[mid+1:], timestamp)
	}
	return pairs[mid].Value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
