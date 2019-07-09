package main

type LRUCache struct {
	capacity   int
	cache      map[int]*listNode
	head, tail *listNode
}

type listNode struct {
	key, value int
	prev, next *listNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &listNode{-1, -1, nil, nil}, &listNode{-1, -1, nil, nil}
	head.next, tail.prev = tail, head
	return LRUCache{capacity, make(map[int]*listNode), head, tail}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.pop(node)
	this.push(node)
	return node.value
}

func (this *LRUCache) pop(node *listNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) push(node *listNode) {
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev.next = node
	this.tail.prev = node
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.cache[key].value = value
		return
	}
	if len(this.cache) == this.capacity {
		first := this.head.next
		this.pop(first)
		delete(this.cache, first.key)
	}
	newNode := &listNode{key, value, nil, nil}
	this.cache[key] = newNode
	this.push(newNode)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
