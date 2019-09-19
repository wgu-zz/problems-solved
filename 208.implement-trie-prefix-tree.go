type Trie struct {
	children [26]*Trie
	isWord   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, c := range word {
		index := byte(c) - 'a'
		if cur.children[index] == nil {
			cur.children[index] = &Trie{}
		}
		cur = cur.children[index]
	}
	cur.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.searchPrefix(word)
	return cur != nil && cur.isWord
}

func (this *Trie) searchPrefix(word string) *Trie {
	cur := this
	for _, c := range word {
		index := byte(c) - 'a'
		if cur.children[index] == nil {
			return nil
		}
		cur = cur.children[index]
	}
	return cur
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
