package Heap

import "data_structure/utils"

type Interface interface {
	utils.Interface
	Push(data interface{})
	Pop() interface{}
}

func Init(h Interface)  {
	for i := 0; i < h.Len(); i++ {
		heapInsert(h, i)
	}
}


func Push(h Interface, data interface{}) {
	h.Push(data)
	heapInsert(h, h.Len()-1)
}

func heapInsert(h Interface, index int) {
	for h.Comparator(index, (index-1)/2) {
		h.Swap(index, (index-1)/2)
		index = (index-1)/2
	}
	//for h.dataSource[index] > h.dataSource[(index-1)/2] {
	//	h.dataSource[index], h.dataSource[(index-1)/2] = h.dataSource[(index-1)/2], h.dataSource[index]
	//	index = (index-1)/2
	//}
}

func heapify(h Interface, index int, heapSize int) {
	left := index * 2 + 1
	var largest int
	for left < heapSize {
		// 两个孩子之间比较出谁大谁小
		//h.dataSource[left] < h.dataSource[left+1]
		if left + 1 < heapSize &&  h.Comparator(left+1, left){
			largest = left + 1
		} else {
			largest = left
		}
		//h.dataSource[index] > h.dataSource[largest]
		if  h.Comparator(index, largest){
			largest = index
		} else {
			largest = largest
		}

		if largest == index {  // 当前位置的值变小后和左右孩子比较后，还是自身的值大时，此时该值不用下沉；你和你孩子的最大值是自己，直接break
			break
		}
		//h.dataSource[largest], h.dataSource[index] = h.dataSource[index], h.dataSource[largest]
		h.Swap(largest, index)
		index = largest        // 下沉
		left = index * 2 + 1
	}
}

func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	data := h.Pop()
	heapSize := h.Len()
	heapify(h, 0, heapSize)

	return data
}

//func (h *ArrayHeap) Peek() interface{} {
//	return h.dataSource.QueryData(0)
//}


type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Comparator(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Push(data interface{}) {
	*h = append(*h, data.(int))
}

func (h *MaxHeap) Pop() interface{} {
	n := h.Len()-1
	data := (*h)[n]
	*h = (*h)[:n]
	return data
}

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Comparator(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Push(data interface{}) {
	*h =  append(*h, data.(int))
}

func (h *MinHeap) Pop() interface{} {
	n := h.Len()-1
	data := (*h)[n]
	*h = (*h)[:n]
	return data
}

