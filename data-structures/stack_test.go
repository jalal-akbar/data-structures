package datastructures

import (
	"fmt"
	"testing"
)

type Stack struct {
	item []int
}

func (s *Stack) Push(i int) {
	s.item = append(s.item, i)
}

func (s *Stack) Pop() int {
	l := len(s.item) - 1
	toRemove := s.item[l]
	s.item = s.item[:l]

	return toRemove
}

func TestStack(t *testing.T) {
	s := &Stack{}
	fmt.Println(s)
	// Push
	s.Push(100)
	s.Push(200)
	s.Push(300)
	fmt.Println(s)
	// Pop
	s.Pop()
	fmt.Println(s)

}
