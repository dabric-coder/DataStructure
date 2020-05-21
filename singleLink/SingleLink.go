package singleLink

import "fmt"

// 节点
type Node struct {
	Data interface{}
	Next *Node
}


// 创建链表
func (node *Node) Create(Data ...interface{}) {
	if node == nil {
		return
	}

	head := node
	for i := 0; i < len(Data); i++ {
		newNode := new(Node)
		newNode.Data = Data[i]
		newNode.Next = nil
		node.Next = newNode
		node = newNode
	}
	node = head
}

// 打印链表
func (node *Node) Print() {
	if node == nil {
		fmt.Println("链表为空")
		return
	}

	cur := node.Next
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Println()

}

// 返回链表长度
func (node *Node) Length() (length int) {
	if node == nil {
		fmt.Println("链表为空")
	}
	cur := node.Next

	for cur != nil {
		length++
		cur = cur.Next
	}
	return
}
//
//
// 插入数据（头插）
func (node *Node) InsertByHead (Data interface{}) {
	if node == nil {
		return
	}
	// 创建一个新的节点，将Data数据写入进去
	newNode := new(Node)
	newNode.Data = Data
	newNode.Next = nil

	// 在链表的头插入元素分两种情况：
	// 1. 链表为空
	// 这种情况下，直接让node.Next指向新的结点
	if node.Next == nil {
		node.Next = newNode
	} else {
		// 2. 链表中有节点
		// 这种情况下，将新的节点插入到node节点和链表中的第一个结点之间
		NextNode := node.Next    // 保留原来第一个节点
		node.Next = newNode		// node节点指向新的节点
		newNode.Next = NextNode  // 新节点指向原先的第一个节点
	}
}
// 插入数据（尾插）
func (node *Node) InsertByTail (Data interface{}) {
	if node == nil {
		return
	}
	// 创建一个新的节点，写入数据
	newNode := &Node{Data, nil}
	/*
		尾插思路：
		如果链表为空(node.Next==nil)，直接将新节点插入到node后，
		如果链表不为空，定义一个节点的指向，利用循环，让其指向链表的最后的一个节点，指向新的节点即可
	*/
	if node.Next == nil {
		node.Next = newNode
	} else {
		cur := node.Next
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newNode
	}
}

//
// 在指定下标位插入数据
func (node *Node) InsertByIndex (index int, Data interface{}) {
	/*
	思路：

	如果该结点为第一个节点，直接调用InsertByTail方法
	如果该结点为链表的最后一个节点，直接调用IndertByHead方法
	最后，找到给出的index位置的前一个节点插入到指定的位置，0 1 2 3 4 5   比如插到 2位置，就是插入到1和2之间
	*/
	if node == nil {
		return
	}
	if index < 0 || index > node.Length() {
		fmt.Println("无效的index")
		return
	}

	// 创建新节点
	newNode := &Node{Data, nil}

	if index == 0 {
		node.InsertByHead(Data)
		return
	}

	cur := node
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	if cur.Next == nil {
		node.InsertByTail(Data)
	} else {
		tmp := cur.Next
		cur.Next = newNode
		newNode.Next = tmp
	}
}
//
// 删除指定下标的数据
func (node *Node) DeleteByIndex(index int) {
	// index从0开始
	/*
	思路：
	如果index==0，只需让node指向链表的第二节点
	找到index对应节点的前一个节点pre
	如果index对应的节点为链表的最后一个节点：
	则让pre指向nil
	其他情况：
	让pre指向index对应节点的下一个节点
	*/
	if node == nil {
		return
	}
	if index < 0 || index > node.Length() {
		fmt.Println("无效index")
		return
	}
	if index == 0 {
		node.Next = node.Next.Next
		return
	}

	pre := node
	for i := 0; i < index; i++ {
		pre = pre.Next
	}

	if pre.Next.Next == nil {
		pre.Next = nil
	} else {
		pre.Next = pre.Next.Next
	}
}
//
// 查找数据（数据）   返回索引
func (node *Node) Search(Data interface{}) (index int) {
	if node == nil {
		return
	}
	cur := node.Next

	for cur != nil {
		if cur.Data == Data {
			return
		}
		cur = cur.Next
		index++
	}
	if cur == nil {
		index = -1   // -1代表索引不存在
	}
	return
}
//
//销毁链表
func (node *Node) Destroy() {
	if node == nil {
		return
	}

	for node != nil {
		node.Data = nil
		node.Next=nil
		node = node.Next
	}
}

