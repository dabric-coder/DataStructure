package ArrayList

type StackArray interface {
	Clear()  // 清空
	Size() int  // 大小
	isEmpty() bool  // 是否满了
	isFull() bool  // 是否为空
	Push(data interface{})  // 压入
	Pop() interface{}  // 弹出
	Peek() interface{}  // 弹出栈顶元素
}

type Stack struct {
	dataSource *ArrayList
	capSize int  // 最大范围
}

func NewStack() *Stack {
	myStack := new(Stack)
	myStack.dataSource = NewArrayList()
	myStack.capSize = 10
	return myStack
}

func (s *Stack) Clear() {
	/*
		golang有其自己的垃圾回收机制。
		重新给dataSource开辟和原始相同的空间，将原始空间覆盖；
		并将capSize与currentSize赋值为最初的值
	*/
	s.dataSource.Clear()
	s.capSize = 10
}
func (s *Stack) Size() int {
	return s.dataSource.theSize
}
func (s *Stack) isEmpty() bool {
	if s.dataSource.theSize == 0 {
		return true
	} else {
		return false
	}
}
func (s *Stack) isFull() bool {
	if s.dataSource.theSize == s.capSize {
		return true
	} else {
		return false
	}
}

func (s *Stack) Push(data interface{})  {
	if !s.isFull() {
		s.dataSource.Append(data)
	}
}

func (s *Stack) Pop() interface{} {
	if !s.isEmpty() {
		data, _ := s.dataSource.Get(s.dataSource.theSize-1)
		s.dataSource.Delete(s.dataSource.theSize-1)
		return data
	}
	return nil
}

func (s *StackX) Peek() interface{} {
	if !s.isEmpty() {
		return s.dataSource.dataStore[s.dataSource.theSize]
	}
	return nil
}
