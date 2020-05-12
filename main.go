package main

import (
	"data_structure/ArrayList"
	"data_structure/Queue/CircleQueue"
	"data_structure/Queue/LinkQueue"
	"data_structure/Stack/LinkStack"
	"data_structure/Stack/StackArray"
	"fmt"
)

func main1() {

	arrayList := ArrayList.NewArrayList()
	arrayList.Append("a1")
	arrayList.Append("a2")
	arrayList.Append("a3")
	arrayList.Append("a4")
	arrayList.Delete(1)
	fmt.Println(arrayList)
}

func main2() {
	// 定义接口对象，赋值的对象必须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	list.Append("d4")
	list.Append("f5")
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		fmt.Println(item)
	}
}

func main3() {
	myStack := StackArray.NewStack()
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
}

func main4() {
	myStack := ArrayList.NewStack()
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
}

func main5() {
	myStack := ArrayList.NewStackX()
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)
	//fmt.Println(myStack.Pop())
	//fmt.Println(myStack.Pop())
	//fmt.Println(myStack.Pop())
	//fmt.Println(myStack.Pop())
	//fmt.Println(myStack.Pop())
	// 使用迭代器遍历栈中的元素
	for it := myStack.MyIt; it.HasNext(); {
		item, _ := it.Next()
		fmt.Println(item)
	}
}

func main6() {
	var myq CircleQueue.CircleQueue
	CircleQueue.InitQueue(&myq)
	for i := 1; i <= 5; i++ {
		_ = CircleQueue.EnQueue(&myq, i)
	}
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
}

func main7() {
	myStack := LinkStack.NewStack()
	fmt.Println(myStack.IsEmpty())

	myStack.Push(1)
	fmt.Println(myStack.IsEmpty())

	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)

	fmt.Println(myStack.Length())

	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Length())

	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
}

func main() {
	myQueue := LinkQueue.NewLinkQueue()
	fmt.Println(myQueue.IsEmpty())
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	myQueue.Enqueue(4)

	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.Dequeue())
}