package Queue

type MyQueue interface {
	Size() int  // 大小
	Front() interface{}  // 返回队首元素
	End() interface{}  // 返回队列中的最后一个元素
	IsEmpty() bool  // 是否为空
	EnQueue(data interface{})  // 入队
	DeQueue() interface{}  // 出队
	Clear()  // 清空
}


type Queue struct {
	dataStore []interface{}  // 队列的数据存储
	theSize int   // 队列的大小
}

func NewQueue() *Queue {
	// 初始化一个Queue
	myQueue := new(Queue)
	myQueue.dataStore = make([]interface{}, 0)
	myQueue.theSize = 0
	return myQueue
}

func (q *Queue) Size() int {
	return q.theSize
}
func (q *Queue) Front() interface{} {
	if !q.IsEmpty() {
		return q.dataStore[0]
	}
	return nil
}
func (q *Queue) End() interface{} {
	if !q.IsEmpty() {
		return q.dataStore[q.theSize-1]
	}
	return nil
}
func (q* Queue) IsEmpty() bool {
	return q.Size() == 0
}
func (q *Queue) EnQueue(data interface{}) {
	q.dataStore = append(q.dataStore, data)
	q.theSize++
}
func (q *Queue) DeQueue() interface{} {
	if q.Size() == 0 {
		return nil
	}
	data := q.dataStore[0]
	if q.Size() > 1 {
		q.dataStore = q.dataStore[1:]
	} else {
		// 当队列有一个元素，此时需要出队，可以重新开辟一个数据存储空间
		q.dataStore = make([]interface{}, 0)
	}
	q.theSize--
	return data
}
func (q *Queue) Clear() {
	q.dataStore = make([]interface{}, 0, 10)
	q.theSize = 0
}