package main

import "testing"

func TestTrie(t *testing.T) {
	commands := []string{"insert", "insert", "insert", "insert", "insert", "insert", "search", "search", "search", "search", "search", "search", "search", "search", "search", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith"}
	values := []string{"app", "apple", "beer", "add", "jam", "rental", "apps", "app", "ad", "applepie", "rest", "jan", "rent", "beer", "jam", "apps", "app", "ad", "applepie", "rest", "jan", "rent", "beer", "jam"}
	expected := []bool{true, true, true, true, true, true, false, true, false, false, false, false, false, true, true, false, true, true, false, false, false, true, true, true}
	trie := Constructor()
	for i, command := range commands {
		switch command {
		case "insert":
			trie.Insert(values[i])
		case "search":
			if expected[i] != trie.Search(values[i]) {
				t.Error("PWNAGE! search: step:", i)
			}
		case "startsWith":
			if expected[i] != trie.StartsWith(values[i]) {
				t.Error("PWNAGE! startsWith: step:", i)
			}
		}
	}

}

type Trie struct {
	Exists bool
	//Rune rune
	Kids map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Exists: false, Kids: make(map[rune]*Trie, int('z')-int('a')+1)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	ptr := this
	for _, r := range word {
		if _, ok := ptr.Kids[r]; !ok {
			ptr.Kids[r] = &Trie{Exists: false, Kids: make(map[rune]*Trie, int('z')-int('a')+1)}
		}
		ptr = ptr.Kids[r]
	}
	ptr.Exists = true
}

func (this *Trie) SearchGen(word string, evaluator func(t *Trie) bool) bool {
	ptr := this
	for _, r := range word {
		if p, exists := ptr.Kids[r]; exists {
			ptr = p
			continue
		}
		return false
	}
	return evaluator(ptr)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.SearchGen(word, func(t *Trie) bool {
		return t.Exists
	})
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchGen(prefix, func(t *Trie) bool {
		return true
	})
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
