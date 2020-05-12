package LinkStack

type Node struct {
	data interface {}
	pNext *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

/*
	链表带头节点
	栈的实现方式：
	头插头删除
*/

func NewStack() *Node {
	return &Node{}  // 返回一个节点指针
}

func (n *Node) IsEmpty() bool {
	if n.pNext == nil {
		return true
	} else {
		return false
	}
}

func (n *Node) Push(data interface{}) {
	newNode := &Node{data, nil}
	newNode.pNext = n.pNext
	n.pNext = newNode  // 头部插入
}

func (n *Node) Pop() interface{} {
	if n.IsEmpty() {
		return nil
	}
	data := n.pNext.data
	n.pNext = n.pNext.pNext
	return data
}

func (n *Node) Length() int {
	cur := n.pNext
	length := 0
	for cur != nil {
		length++
		cur = cur.pNext
	}
	return length
}