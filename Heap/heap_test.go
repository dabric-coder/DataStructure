package Heap

import (
	"fmt"
	"testing"
)

func TestArrayHeap(t *testing.T) {
	heap := new(MaxHeap)
	heap.Push(3)
	heap.Push(4)
	heap.Push(5)
	heap.Push(0)
	heap.Push(1)
	heap.Push(6)

	Init(heap)
	fmt.Println(*heap)

	Push(heap, 7)
	fmt.Println(*heap)

	fmt.Println(Pop(heap))
	fmt.Println(*heap)
}
