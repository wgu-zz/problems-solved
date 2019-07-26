package main

import "math/rand"

type RandomizedSet struct {
	set map[int]int
	a   []int
}

/** Initialize your data structure here. */
func Constructor2() RandomizedSet {
	return RandomizedSet{make(map[int]int), []int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}
	this.set[val] = len(this.a)
	this.a = append(this.a, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.set[val]
	if !ok {
		return false
	}
	last := len(this.a) - 1
	this.set[this.a[last]] = index
	this.a[index] = this.a[last]
	this.a = this.a[:last]
	delete(this.set, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.a))
	return this.a[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
