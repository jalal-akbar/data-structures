package datastructures

import (
	"fmt"
	"testing"
)

type Quee struct {
	quee []int
}

// Enquee
func (q *Quee) Enquee(value int) {
	q.quee = append(q.quee, value)
}

// Dequee
func (q *Quee) Dequee() int {
	toRemove := q.quee[0]
	q.quee = q.quee[1:]
	return toRemove
}

func TestQuee(t *testing.T) {
	quee := &Quee{}
	quee.Enquee(40)
	quee.Enquee(20)
	quee.Enquee(30)
	fmt.Println(quee)
	deq1 := quee.Dequee()
	deq2 := quee.Dequee()
	fmt.Println(deq1)
	fmt.Println(deq2)
}
