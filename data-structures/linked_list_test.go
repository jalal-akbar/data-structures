package datastructures

import (
	"fmt"
	"testing"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(n *Node) {
	oldHead := l.Head     // Head Before Insert
	l.Head = n            // Insert Data
	l.Head.Next = oldHead // link oldHead to second of newHead
	l.Length++            // increament length linkedList
}

func (l LinkedList) ToPrint() {
	toPrint := l.Head
	len := l.Length
	for len != 0 {
		fmt.Printf("%v ", toPrint.Data)
		toPrint = toPrint.Next
		len--
	}
	fmt.Printf("\n")
}

func (l *LinkedList) ToDelete(value int) {
	// if l.Length == 0 {
	// 	return
	// }
	// Delete Head Data
	if l.Head.Data == value {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	toDelete := l.Head
	for toDelete.Next.Data != value {
		if toDelete.Next.Next == nil { // when data doesn't exist
			return
		}
		toDelete = toDelete.Next
	}
	toDelete.Next = toDelete.Next.Next
	l.Length--
}

func TestLinkedList(t *testing.T) {
	myList := &LinkedList{}
	node1 := &Node{Data: 10}
	node2 := &Node{Data: 20}
	node3 := &Node{Data: 30}
	node4 := &Node{Data: 40}
	node5 := &Node{Data: 50}
	// Prepend
	myList.Prepend(node1)
	myList.Prepend(node2)
	myList.Prepend(node3)
	myList.Prepend(node4)
	myList.Prepend(node5)

	myList.ToPrint()
	fmt.Println("All myList :", &myList)

	//Delete
	myList.ToDelete(30) // Delete Existing Value
	myList.ToPrint()
	// myList.ToDelete(60) // Delete Not Existing Value
	// myList.ToPrint()
	// myList.ToDelete(50) // Delete Head
	// myList.ToPrint()

}
