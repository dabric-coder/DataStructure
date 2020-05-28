package BinarySearchTree

import (
	"DataStructure/Queue"
	"DataStructure/Stack/StackArray"
	"bytes"
	"fmt"
	"math"
	"strconv"
)

type BST struct {
	root       *Node
	size       int
	comparable Comparable
}

// 构造一个二叉树
func NewBST(comparable Comparable) *BST {
	return &BST{nil, 0, comparable}
}

func (bst *BST) Size() int {
	return bst.size
}

func (bst *BST) IsEmpty() bool {
	return bst.size == 0
}

func (bst *BST) Clear() {
	bst.root = nil
	bst.size = 0
}

// 添加元素
func (bst *BST) Add(data interface{}) {
	if bst.root == nil {
		bst.root = &Node{data, nil, nil, nil}
		bst.size++
		return
	}

	parent := bst.root
	node := bst.root
	cmp := 0

	for node != nil {
		cmp = bst.comparable(data, node.data)
		parent = node
		if cmp > 0 {
			node = node.right
		} else if cmp < 0 {
			node = node.left
		} else {
			node.data = data
			return
		}
	}

	// 循环结束找到要添加位置的父节点，此时还要判断添加到父节点的左子树还是右子树
	newNode := &Node{data, nil, nil, parent}
	if cmp > 0 {
		parent.right = newNode
	} else {
		parent.left = newNode
	}
	bst.size++
}

// 前序遍历
func (bst *BST) PreOrderTraversal() {
	bst.preOrderTraversal(bst.root)
}

func (bst *BST) preOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	//fmt.Printf("%v ", root.data)
	//bst.preOrderTraversal(root.left)
	//bst.preOrderTraversal(root.right)
	s := StackArray.NewStack()
	s.Push(root)

	for !s.IsEmpty() {
		cur := s.Pop().(*Node)
		fmt.Printf("%v ", cur.data)
		if cur.right != nil {
			s.Push(cur.right)
		}
		if cur.left != nil {
			s.Push(cur.left)
		}
	}
}

// 中序遍历
func (bst *BST) InOrderTraversal() {
	bst.inOrderTraversal(bst.root)
}

func (bst *BST) inOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	//bst.inOrderTraversal(root.left)
	//fmt.Printf("%v ", root.data)
	//bst.inOrderTraversal(root.right)
	s := StackArray.NewStack()
	cur := root

	for cur != nil || !s.IsEmpty() {
		if cur != nil {
			s.Push(cur)
			cur = cur.left
		}else {
			cur = s.Pop().(*Node)
			fmt.Printf("%v ", cur.data)
			cur = cur.right
		}
	}
}

// 后序遍历
func (bst *BST) PostOrderTraversal() {
	if bst.root == nil {
		return
	}
	bst.postOrderTraversal(bst.root)
}

func (bst *BST) postOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	// 递归实现后序遍历
	//bst.postOrderTraversal(root.left)
	//bst.postOrderTraversal(root.right)
	//fmt.Printf("%v ", root.data)

	// 非递归实现后序遍历
	s := StackArray.NewStack()

	lastVisitNode := new(Node)
	cur := root
	s.Push(root)

	// cur代表栈顶元素
	for !s.IsEmpty() && cur != nil {
		if (cur.left == nil && cur.right == nil) || (cur.left == lastVisitNode || cur.right == lastVisitNode) {
			node := s.Pop().(*Node)
			lastVisitNode = node
			fmt.Printf("%v ", node.data)

			if s.Peek() != nil {
				cur = s.Peek().(*Node)
			}
		} else {
			if cur.right != nil {
				s.Push(cur.right)
			}

			if cur.left != nil {
				s.Push(cur.left)
			}
			if s.Peek() != nil {
				cur = s.Peek().(*Node)
			}
		}

	}
}

// 层序遍历
func (bst *BST) LevelOrderTraversal() {
	bst.levelOrderTraversal(bst.root)
}

func (bst *BST) levelOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	q := Queue.NewQueue()
	q.EnQueue(root)


	for !q.IsEmpty() {
		tmpNode := q.DeQueue()
		node := tmpNode.(*Node)
		fmt.Printf("%v ", node.data)

		if node.left != nil {
			q.EnQueue(node.left)
		}

		if node.right != nil{
			q.EnQueue(node.right)
		}
	}

}


func (bst *BST) Height2() int {
	return bst.height(bst.root)
}

func (bst *BST) height(node *Node) int {
	if node == nil {
		return 0
	}
	maxChildHeight := int(math.Max(float64(bst.height(node.left)), float64(bst.height(node.right))))
	return  maxChildHeight + 1
}

func (bst *BST) Height3() int {
	if bst.root == nil {
		return 0
	}

	// 树的高度
	height := 0
	levelSize := 1  // 每一层节点个数

	q := Queue.NewQueue()
	q.EnQueue(bst.root)

	for !q.IsEmpty() {

		node := q.DeQueue().(*Node)
		levelSize--

		if node.left != nil {
			q.EnQueue(node.left)
		}

		if node.right != nil {
			q.EnQueue(node.right)
		}

		if levelSize == 0 {  // 意味着将要访问下一层的节点
			levelSize = q.Size()
			height++
		}
	}
	return height
}


func (bst *BST) Height() int {
	if bst.root == nil {
		return 0
	}

	// 树的高度
	height := 0

	q := Queue.NewQueue()
	q.EnQueue(bst.root)

	for !q.IsEmpty() {
		levelSize := q.Size()  // 每层的节点数
		for i := 0; i < levelSize; i++ {
			node := q.DeQueue().(*Node)
			if node.left != nil {
				q.EnQueue(node.left)
			}

			if node.right != nil {
				q.EnQueue(node.right)
			}

			if levelSize == 0 {  // 意味着将要访问下一层的节点
				levelSize = q.Size()
				height++
			}
		}
		height++
	}
	return height
}

func (bst *BST) isComplete() bool {
	if bst.root == nil {
		return false
	}

	q := Queue.NewQueue()
	q.EnQueue(bst.root)

	leaf := false

	for !q.IsEmpty() {
		node := q.DeQueue().(*Node)
		if leaf && (node.left != nil || node.right != nil) {
			return false
		}
		//if (leaf && (node.left != nil || node.right != nil)) || (node.left == nil && node.right != nil ) {
		//	return false
		//}
		//
		//if node.left != nil {
		//	q.EnQueue(node.left)
		//}
		//
		//if node.right != nil {
		//	q.EnQueue(node.right)
		//} else {
		//	leaf = true
		//}

		if node.left != nil && node.right != nil {
			q.EnQueue(node.left)
			q.EnQueue(node.right)
		} else if node.left == nil && node.right != nil {
			return false
		} else {
			leaf = true
		}
	}
	return true
}

func (bst *BST) Predecessor(node *Node) *Node {
	if node == nil {
		return nil
	}

	p := node.left

	if p != nil {
		for p.right != nil {
			p = p.right
		}
		return p
	}

	for node.parent != nil && node == node.parent.left {
		node = node.parent
	}
	return node.parent
}

func (bst *BST) PredecessorWithStack(node *Node) *Node {

	// 判断node节点是否有左子树，如果有左子树找到左子树最右边的节点，就为node的前驱节点；
	// 如果没有左子树有右子树，使用栈结构记录从root到node节点的路径
	// 那么栈中存储的是node的父节点的父节点的父一直到root
	// 往上找，一直找到当前节点是父节点的右孩子，那么该父节点为node的前驱节点

	p := node.left

	// 往下找
	if p != nil {
		for p.right != nil {
			p = p.right
		}
		return p
	}

	// 往上找

	cur := bst.root

	myStack := make([]*Node, 0)

	for cur != node {
		myStack = append(myStack, cur)
		if  bst.comparable(node.data, cur.data) > 0 {
			cur = cur.right
		} else {
			cur = cur.left
		}
	}

	for len(myStack) > 0 && myStack[len(myStack)-1].right != node {
		node = myStack[len(myStack)-1]
		myStack = myStack[:len(myStack)-1]
	}

	if len(myStack) > 0 {
		return myStack[len(myStack)-1]
	}
	return nil

}

func (bst *BST) Successor(node *Node) *Node {

	if node == nil {
		return nil
	}

	p := node.right

	if p != nil {
		for p.left != nil {
			p = p.left
		}
		return p
	}

	for node.parent != nil && node == node.parent.right {
		node = node.parent
	}
	return node.parent
}


func (bst *BST) SuccessorWithStack(node *Node) *Node {
	if node == nil {
		return nil
	}

	// 当前节点有右子树，向下找
	p := node.right

	if p != nil {
		for p.left != nil {
			p = p.left
		}
		return p
	}

	// 当前节点没有右子树，向上找

	// 首先使用栈记录从root到node的路径

	myStack := make([]*Node, 0)

	cur := bst.root

	for cur != node {
		myStack = append(myStack, cur)
		if bst.comparable(node.data, cur.data) > 0 {
			cur = cur.right
		} else {
			cur = cur.left
		}
	}

	for len(myStack) != 0 && myStack[len(myStack)-1].left != node {
		node = myStack[len(myStack)-1]
		myStack = myStack[0:len(myStack)-1]
	}
	if len(myStack) > 0 {
		return myStack[len(myStack)-1]
	}
	return nil
}


func (bst *BST) Remove(node *Node) {
	if node == nil {
		return
	}
	bst.size--

	if node.left != nil && node.right != nil { // 节点的度为2
		// 找出该节点的后继节点，将后继节点的值赋值个该节点，然后删除
		s := bst.Successor(node)
		node.data = s.data
		node = s
	}

	// 要删除的节点node的度要么为1要么为0

	var replacement *Node
	if node.left != nil {
		replacement = node.left
	} else {
		replacement = node.right
	}

	if replacement != nil { // 表示要删除的节点的度为1
		replacement.parent = node.parent
		if node.parent == nil { // node的度为1并且为根节点
			bst.root = replacement
		} else if node == node.parent.left {
			node.parent.left = replacement
		} else {
			node.parent.right = replacement
		}

	} else if node.parent == nil { // node是叶子节点并且是根节点
		bst.root = nil

	} else {
		// node是叶子节点但不是根节点
		if node == node.parent.left {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
	}
}

// 打印二叉树
func (bst *BST) String() string {
	var buffer bytes.Buffer   // 保存字符串
	bst.GenerateBSTString(bst.root, 0, "H", 17)
	return buffer.String()
}

func (bst *BST) GenerateBSTString(root *Node, depth int, to string, length int) {
	if root == nil {
		return
	}
	bst.GenerateBSTString(root.right, depth + 1, "v", length)
	val := to + strconv.Itoa(root.data.(int)) + to
	lenM := len(val)
	lenL := (length - lenM) / 2
	lenR := length - lenM - lenL
	val = bst.GenerateDepthString(lenL) + val + bst.GenerateDepthString(lenR)
	fmt.Println(bst.GenerateDepthString(depth * length) + val)
	bst.GenerateBSTString(root.left, depth+1, "^", length)
}

func (bst *BST) GenerateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString(" ")
	}
	return buffer.String()
}

