type LRUCache struct {
	cache      map[int]*Node
	head, tail *Node
	capacity   int
}

type Node struct {
	key, value int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{make(map[int]*Node), head, tail, capacity}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.remove(n)
	this.insert(n)
	return n.value
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.cache[key].value = value
		return
	}
	if len(this.cache) == this.capacity {
		n := this.head.next
		delete(this.cache, n.key)
		this.remove(n)
	}
	n := &Node{key: key, value: value}
	this.cache[key] = n
	this.insert(n)
}

func (this *LRUCache) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev, n.next = nil, nil
}

func (this *LRUCache) insert(n *Node) {
	this.tail.prev.next = n
	n.prev = this.tail.prev
	this.tail.prev = n
	n.next = this.tail
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
