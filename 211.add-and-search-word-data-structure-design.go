type WordDictionary struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	root := &Node{false, [26]*Node{}}
	return WordDictionary{root}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for index, c := range word {
		i := byte(c) - 'a'
		if cur.children[i] == nil {
			cur.children[i] = &Node{false, [26]*Node{}}
		}
		cur = cur.children[i]
		if index == len(word)-1 {
			cur.isWord = true
		}
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return searchSub(word, this.root)
}

func searchSub(word string, cur *Node) bool {
	for i, c := range word {
		if c == '.' {
			for _, n := range cur.children {
				if n != nil && searchSub(word[i+1:], n) {
					return true
				}
			}
			return false
		} else {
			i := byte(c) - 'a'
			if cur.children[i] == nil {
				return false
			}
			cur = cur.children[i]
		}
	}
	return cur.isWord
}

type Node struct {
	isWord   bool
	children [26]*Node
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
