package main

// The key of map is standing the char of the word.
type Trie struct {
	IsEnd bool
	Next  map[byte]*Trie
}

func TrieConstructor() Trie {
	return Trie{
		IsEnd: false,
		Next:  make(map[byte]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		if node, ok := cur.Next[word[i]]; ok {
			cur = node
		} else {
			node = &Trie{
				Next: make(map[byte]*Trie),
			}
			cur.Next[word[i]] = node
			cur = node
		}
	}
	cur.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		cur = cur.Next[word[i]]
		if cur == nil {
			return false
		}
	}
	// When end loop, cur is the end char of the word(like '\0' in C language).
	return cur.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		cur = cur.Next[prefix[i]]
		if cur == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
