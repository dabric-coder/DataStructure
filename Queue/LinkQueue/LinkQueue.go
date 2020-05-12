package LinkQueue

type Node struct {
	data interface{}
	pNext *Node
}

type QueueLink struct {
	front *Node
	rear *Node
}

type LinkQueue interface {
	IsEmpty() bool
	Enqueue(data interface{})
	Dequeue() interface{}
	Length() int
}

/*
	链式队列有两种方式：
	头部插入，尾部删除
	头部删除，尾部追加
*/

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (q *QueueLink) IsEmpty() bool {
	if q.front == nil {
		return true
	} else {
		return false
	}
}

func (q *QueueLink) Enqueue(data interface{}) {
	newNode := &Node{data, nil}
	if q.IsEmpty() {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.pNext = newNode
		q.rear = q.rear.pNext
	}
}

func (q *QueueLink) Dequeue() interface{} {
	// 头部删除
	if q.IsEmpty() {
		return nil
	}
	data := q.front.data  // 记录头部位置
	if q.front == q.rear {  // 只剩一个
		q.front = nil
		q.rear = nil
	} else {
		q.front = q.front.pNext
	}
	return data
}

func (q *QueueLink) Length() int {
	cur := q.front
	length := 0
	for cur != nil {
		length++
		cur = cur.pNext
	}
	return length
}


