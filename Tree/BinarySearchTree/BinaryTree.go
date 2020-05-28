package BinarySearchTree


type BinaryTree interface {
	Size() int   // 元素的数量
	IsEmpty() bool  // 是否为空
	Clear()  // 清空所有元素
	Add(data interface{})  // 添加元素
	Remove(data interface{})  // 删除元素
	Contains(data interface{})  // 是否包含某元素
}

type Comparable func (i, j interface{}) int  // 比较器

type Node struct {
	data interface{}
	left *Node
	right * Node
	parent *Node
}