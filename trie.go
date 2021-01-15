package trie

import "fmt"

type TrieNode struct {
	children map[string]*TrieNode
	value    *int
}

type Trie struct {
	root *TrieNode
}

func New() Trie {
	t := Trie{}
	t.root = newNode()

	return t
}

func (t *TrieNode) setValue(value int) {
	(t.value) = &value
}

func (t Trie) Display() {
	fmt.Println(t.root.children)
}

func newNode() *TrieNode {
	hash := make(map[string]*TrieNode)
	return &TrieNode{hash, nil}
}

func (t Trie) Insert(word string, value int) {
	var currentNode *TrieNode = t.root

	for idx := range word {
		char := string(word[idx])
		if contains(currentNode, char) {
			currentNode = currentNode.children[char]
		} else {
			newNode := newNode()
			currentNode.children[char] = newNode
			currentNode = newNode
		}
	}

	currentNode.children["*"] = nil
	currentNode.value = &value
}

func (t Trie) Search(word string) *TrieNode {
	currentNode := t.root

	for idx := range word {
		char := string(word[idx])
		if contains(currentNode, char) {
			currentNode = currentNode.children[char]

		} else {
			return nil
		}
	}

	return currentNode
}

func (t *TrieNode) collectAllWords(word string, maxWord *string, maxValue *int, firstCall bool) *string {
	for char, child := range t.children {
		if char == "*" {
			// *words = append(*words, word)

			if *(t.value) > *maxValue && firstCall != true {
				*maxValue = *t.value
				*maxWord = word
			}
		} else {
			child.collectAllWords(word+char, maxWord, maxValue, false)
		}
	}

	return maxWord
}

func contains(hash *TrieNode, value string) bool {
	hashVal := hash.children[value]

	valType := fmt.Sprintf("%v", hashVal)

	if valType != "<nil>" {
		return true
	}

	return false
}

func (t Trie) Autocomplete(prefix string) string {

	node := t.Search(prefix)

	if node == nil {
		return ""
	}

	var word string = ""
	var max int = 0
	return *(node.collectAllWords("", &word, &max, true))
}
