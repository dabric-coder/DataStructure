package Heap

import (
	"fmt"
	"testing"
)

type Person struct {
	age int
	name string
}

func maxHeapComparator (i, j interface{}) int {
	return i.(int) - j.(int)
}

func minHeapComparator (i, j interface{}) int {
	return j.(Person).age - i.(Person).age
}

func TestArrayHeap(t *testing.T) {
	arr := []interface{}{3,4,5,0,1,6}
	MyHeap := NewHeapifyBinaryHeap(arr, maxHeapComparator)
	fmt.Println(MyHeap.data)
	MyHeap.heapify()

	fmt.Println(MyHeap.data)

	MyHeap.add(2)

	fmt.Println(MyHeap.data)

	//fmt.Println(MyHeap.remove())
	//fmt.Println(MyHeap.data)
	//MyHeap.replace(10)
	//fmt.Println(MyHeap.data)
}


//func TestPersonHeap(t *testing.T) {
//	person1 := Person{30, "tom"}
//	person2 := Person{40, "tom"}
//	person3 := Person{50, "tom"}
//	person4 := Person{10, "tom"}
//	person5 := Person{60, "tom"}
//
//	MyHeap := NewBinaryHeap(minHeapComparator)
//
//	MyHeap.add(person1)
//	MyHeap.add(person2)
//	MyHeap.add(person3)
//	MyHeap.add(person4)
//	MyHeap.add(person5)
//
//	fmt.Println(MyHeap.data)
//
//}