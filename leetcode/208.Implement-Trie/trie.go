package leetcode

type Trie struct {
	trieNode [26]*Trie
	end      bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		index := v - 'a'
		if node.trieNode[index] == nil {
			node.trieNode[index] = &Trie{}
		}
		node = node.trieNode[index]
	}
	node.end = true
}

func (this *Trie) searchPrefix(word string) *Trie {
	node := this
	for _, v := range word {
		index := v - 'a'
		if node.trieNode[index] == nil {
			return nil
		} else {
			node = node.trieNode[index]
		}
	}
	return node
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, v := range word {
		index := v - 'a'
		if node.trieNode[index] == nil {
			return false
		} else {
			node = node.trieNode[index]
		}
	}
	return node.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, v := range prefix {
		index := v - 'a'
		if node.trieNode[index] == nil {
			return false
		} else {
			node = node.trieNode[index]
		}
	}
	return true
}
