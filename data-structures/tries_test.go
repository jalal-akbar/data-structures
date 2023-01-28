package datastructures

import (
	"fmt"
	"testing"
)

const alphabetSize = 26

// nodeTrie represent each node in the trie
type nodeTrie struct {
	Children [alphabetSize]*nodeTrie
	IsEnd    bool
}

// trie represent a trie and has a pointer to the root node
type trie struct {
	Root *nodeTrie
}

// initTrie will create new trie
func initTrie() *trie {
	res := &trie{Root: &nodeTrie{}}
	return res
}

// Insert will take a word and add it to the trie
func (t *trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.Root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			currentNode.Children[charIndex] = &nodeTrie{}
		}
		currentNode = currentNode.Children[charIndex]
	}
	currentNode.IsEnd = true
}

// Seacrh will take in a word and RRETURN true is that word included in the trie
func (t *trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.Root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			currentNode.Children[charIndex] = &nodeTrie{}
		}
		currentNode = currentNode.Children[charIndex]
	}
	return currentNode.IsEnd == true
}
func TestTries(t *testing.T) {
	myTrie := initTrie()
	myTrie.Insert("oreo")
	myTrie.Insert("orc")
	fmt.Println(myTrie.Search("orc"))
}
