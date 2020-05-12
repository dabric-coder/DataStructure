package StackArray

import (
	"fmt"
)

type StackArray interface {
	Clear()  // 清空
	Size() int  // 大小
	IsEmpty() bool  // 是否满了
	IsFull() bool  // 是否为空
	Push(data interface{})  // 压入
	Pop() interface{}  // 弹出
	Peek() interface{}  // 弹出栈顶元素
}

type Stack struct {
	dataSource []interface{}
	capSize int  // 最大范围
	currentSize int  // 实际使用大小
}

func NewStack() *Stack {
	myStack := new(Stack)
	myStack.dataSource = make([]interface{}, 0, 100)
	myStack.capSize = 100
	myStack.currentSize = 0
	return myStack
}

func (s *Stack) Clear() {
	/*
		golang有其自己的垃圾回收机制。
		重新给dataSource开辟和原始相同的空间，将原始空间覆盖；
		并将capSize与currentSize赋值为最初的值
	*/
	s.dataSource = make([]interface{}, 0, 100)
	s.capSize = 100
	s.currentSize = 0
}
func (s *Stack) Size() int {
	return s.currentSize
}
func (s *Stack) IsEmpty() bool {
	if s.currentSize == 0 {
		return true
	} else {
		return false
	}
}
func (s *Stack) IsFull() bool {
	if s.currentSize == s.capSize {
		return true
	} else {
		return false
	}
}

func (s *Stack) Push(data interface{})  {
	if !s.IsFull() {
		s.dataSource = append(s.dataSource, data)
		s.currentSize++
	}
}

func (s *Stack) Pop() interface{} {
	if !s.IsEmpty() {
		data := s.dataSource[s.currentSize-1]
		s.dataSource = s.dataSource[:s.currentSize-1]
		s.currentSize--
		return data
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if !s.IsEmpty() {
		return s.dataSource[s.currentSize-1]
	}
	return nil
}

func (s *Stack) printStack() {
	fmt.Printf("stack: [")
	for i := 0; i < s.currentSize; i++ {
		fmt.Printf("%d ", s.dataSource[i])
	}
	fmt.Println("] top")
}

// 重写栈打印时的展示形式，只需要重新String方法
//func (s *Stack) String() string {
//	var buffer bytes.Buffer
//	buffer.WriteString(fmt.Sprintf(""))
//	fmt.Println()
//
//}
