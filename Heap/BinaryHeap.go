package Heap

import (
	"fmt"
)

type Comparator func(i, j interface{}) int


type BinaryHeap struct {
	data []interface{}
	size int
	comparator Comparator
}

func NewBinaryHeap(comparator Comparator) *BinaryHeap {
	return &BinaryHeap{make([]interface{}, 0), 0, comparator}
}

func NewHeapifyBinaryHeap(data []interface{}, comparator Comparator) * BinaryHeap {
	return &BinaryHeap{data, len(data),comparator}
}

func (BH *BinaryHeap) length() int {
	return BH.size
}

func (BH *BinaryHeap) isEmpty() bool {
	if BH.size == 0 {
		return true
	}
	return false
}

func (BH *BinaryHeap) clear() {
	BH.size = 0
	BH.data = make([]interface{}, 0)
}

func (BH *BinaryHeap) get() interface{} {
	if BH.isEmpty() {
		fmt.Println("heap is empty")
		return nil
	}
	return BH.data[0]
}

func (BH *BinaryHeap) siftUp(index int) {

	//for BH.comparator(BH.data[index], BH.data[(index-1)/2]) {
	//	BH.data[index], BH.data[(index-1)/2] = BH.data[(index-1)/2], BH.data[index]
	//	index = (index-1) /2
	//}

	data := BH.data[index]

	for index > 0 {
		parentIndex := (index-1) >> 1
		parent := BH.data[parentIndex]

		if BH.comparator(data, parent) <= 0 {
			break
		}
		BH.data[index] = parent
		index = parentIndex
	}
	BH.data[index] = data
}

func (BH *BinaryHeap) siftDown(index int) {
	leftChild := index * 2 + 1
	var largetst int

	for leftChild < BH.size {

		if leftChild+1 < BH.size && BH.comparator(BH.data[leftChild+1], BH.data[leftChild]) > 0 {
			largetst = leftChild + 1
		} else {
			largetst = leftChild
		}

		if BH.comparator(BH.data[index], BH.data[largetst]) > 0 {
			largetst = index
		}

		if index == largetst {
			break
		}

		BH.data[index], BH.data[largetst] = BH.data[largetst], BH.data[index]
		index = largetst
		leftChild = index * 2 + 1
	}
}

func (BH *BinaryHeap) add(data interface{}) {
	BH.data = append(BH.data, data)
	BH.size++
	BH.siftUp(BH.size-1)
}

func (BH *BinaryHeap) remove() interface{} {
	if BH.isEmpty() {
		fmt.Println("堆为空")
		return nil
	}
	data := BH.data[0]
	BH.data[0] = BH.data[BH.size-1]
	BH.data = BH.data[0:BH.size-1]
	BH.size--

	BH.siftDown(0)

	return data
}

func (BH *BinaryHeap) replace(data interface{}) interface{} {
	// 删除堆顶元素，同时插入一个新的元素
	var result interface{}
	if BH.isEmpty() {
		BH.data[0] = data
		BH.size++
	} else {
		result = BH.data[0]
		BH.data[0] = data
		BH.siftDown(0)
	}
	return result
}

func (BH *BinaryHeap) heapify() {
	//for i := 1; i < BH.size; i++ {
	//	BH.siftUp(i)
	//}

	for i := BH.size / 2 - 1; i >= 0; i-- {
		// 其中BH.size/2为完全二叉树中的第一个叶子节点
		BH.siftDown(i)
	}
}

