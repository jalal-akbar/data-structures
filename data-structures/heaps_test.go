package datastructures

import "fmt"

// MaxHeaps struct has a slice that holds the array
type MaxHeaps struct {
	array []int
}

// Method

// Insert add element to the heap
func (h *MaxHeaps) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapyUp(len(h.array) - 1)
}

// Extract return the largest key, and removes it from heap

func (h *MaxHeaps) Extract() int {
	extracted := h.array[0]
	last := len(h.array) - 1
	// when array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}
	// take out the last index and put in the root
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.maxHeapyDown(0)
	return extracted
}

// maxHeapyDown will heapy bottom to top
func (h *MaxHeaps) maxHeapyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapyDown will heapy top to bottom
func (h *MaxHeaps) maxHeapyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only one child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeaps) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]

}

// Function

func parent(pa int) int {
	return (pa - 1) / 2
}

func left(le int) int {
	// odd number
	return 2*le + 1
}

func right(ri int) int {
	// even number
	return 2*ri + 2
}
