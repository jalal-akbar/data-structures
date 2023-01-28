package datastructures

import (
	"fmt"
	"testing"
)

var count int

// node represent components of binary tree search
type node struct {
	Key   int
	Left  *node
	Right *node
}

// insert will add node to the tree
// the key to add should not be ready in the tree
func (tree *node) insert(value int) {
	if tree.Key < value {
		// move right
		if tree.Right == nil {
			tree.Right = &node{Key: value}
		} else {
			tree.Right.insert(value)
		}
	} else if tree.Key > value {
		// move left
		if tree.Left == nil {
			tree.Left = &node{Key: value}
		} else {
			tree.Left.insert(value)
		}
	}
}

// search will take in a key value
// and RETURN true if there is a node with that value
func (n *node) search(k int) bool {
	count++
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.search(k)
	} else if n.Key > k {
		return n.Left.search(k)
	}
	return true
}

func TestBinarySearchTree(t *testing.T) {
	tree := &node{Key: 500}
	// left
	tree.insert(450)
	tree.insert(475)
	tree.insert(425)
	tree.insert(415)
	tree.insert(485)
	// right
	tree.insert(550)
	tree.insert(525)
	tree.insert(535)
	tree.insert(515)
	tree.insert(600)
	tree.insert(650)
	tree.insert(575)

	fmt.Println(tree)

	fmt.Println(tree.search(650))
	fmt.Println(count)
}
