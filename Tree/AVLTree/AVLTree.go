package AVLTree

import (
	"bytes"
	"fmt"
	"strconv"
)

type Comparable func (i, j interface{}) int

type AVLNode struct {
	data interface{}
	left *AVLNode
	right *AVLNode
	height int
	parent *AVLNode
}

type AVLTree struct {
	root *AVLNode
	size int
	comparable Comparable
}


func NewAVLNode(comparable Comparable) *AVLTree {
	return &AVLTree{nil,0, comparable}
}


func (avl *AVLTree) Add(data interface{}) {
	if avl.root == nil {
		avl.root = &AVLNode{data, nil, nil, 0, nil}
		avl.size++
		return
	}

	// 如果根节点存在，那么就要看该节点应当存在左子树还是存到右子树
	cur := avl.root

	// 找到要添加的位置的父节点
	parent := avl.root
	cmp := 0

	for cur != nil {
		cmp = avl.comparable(data, cur.data)
		parent = cur
		if cmp > 0 {
			cur = cur.right
		} else if cmp < 0 {
			cur = cur.left
		} else {
			cur.data = data
			return
		}
	}

	newNode := &AVLNode{data, nil, nil, 0, parent}

	if cmp > 0 {
		parent.right = newNode
	} else {
		parent.left = newNode
	}
	avl.size++
}

// 二叉树的遍历

func (avl *AVLTree) PreOrderTraversal() {
	// 使用迭代法遍历
	/*
	借用一个栈结构
	如果root不为空，将root压入栈中，如果栈不为空开始循环：
	弹出栈顶元素，为cur节点，如果其存在右孩子，将右孩子压入栈中
	如果存在左孩子将左孩子压入栈中
	*/
	myStack := make([]*AVLNode, 0)
	if avl.root == nil {
		return
	}

	cur := avl.root
	myStack= append(myStack, cur)  // 将根节点压入栈中

	for len(myStack) > 0 {
		// 弹出栈顶元素
		cur = myStack[len(myStack)-1]
		myStack = myStack[:len(myStack)-1]
		fmt.Printf("%v ", cur.data)

		if cur.right != nil {
			myStack = append(myStack, cur.right)
		}

		if cur.left != nil {
			myStack = append(myStack, cur.left)
		}
	}
}

func (avl *AVLTree) InOrderTraversal() {
	/*
	使用栈结构
	首先当前结点为cur = root
	如果栈不为空或者cur节点不为空，开始循环
	如果cur不为空，将cur压入栈中，cur往左
	如果cur为空，弹出栈顶元素，cur往右跑
	*/
	myStack := make([]*AVLNode, 0)
	cur := avl.root

	for len(myStack) > 0 || cur != nil {
		if cur != nil {
			myStack = append(myStack, cur)
			cur = cur.left
		} else {
			cur = myStack[len(myStack)-1]
			myStack = myStack[:len(myStack)-1]
			fmt.Printf("%v ", cur.data)
			cur = cur.right
		}
	}
}

func (avl *AVLTree) PostOrderTraversal() {
	/*
	 root压入栈中，栈不为空开始循环
	如果栈顶元素是叶子节点，或者上一次访问的节点是栈顶元素的孩子节点，弹出栈顶元素，进行访问
	否则，将栈顶元素的右孩子和左孩子分别压入栈中

	*/
	if avl.root == nil {
		return
	}

	myStack := make([]*AVLNode, 0)
	myStack = append(myStack, avl.root)
	lastVisitNode := new(AVLNode)

	for len(myStack) > 0 {
		cur := myStack[len(myStack)-1]  // 栈顶元素

		if (cur.left == nil && cur.right == nil) || (cur.left == lastVisitNode || cur.right == lastVisitNode) {
			node := myStack[len(myStack)-1]
			myStack = myStack[:len(myStack)-1]
			fmt.Printf("%v ", node.data)
			lastVisitNode = node   // 更新访问节点
		} else {
			if cur.right != nil {
				myStack = append(myStack, cur.right)
			}

			if cur.left != nil {
				myStack = append(myStack, cur.left)
			}
		}
	}
}


func (avl *AVLTree)  Predecessor(node *AVLNode) *AVLNode {
	node = node.left

	for node.right != nil {
		node = node.right
	}
	return node
}



func (avl *AVLTree) String() string {
	var buffer bytes.Buffer   // 保存字符串
	avl.GenerateBSTString(avl.root, 0, "H", 17)
	return buffer.String()
}

func (avl *AVLTree) GenerateBSTString(root *AVLNode, depth int, to string, length int) {
	if root == nil {
		return
	}
	avl.GenerateBSTString(root.right, depth + 1, "v", length)
	val := to + strconv.Itoa(root.data.(int)) + to
	lenM := len(val)
	lenL := (length - lenM) / 2
	lenR := length - lenM - lenL
	val =avl.GenerateDepthString(lenL) + val + avl.GenerateDepthString(lenR)
	fmt.Println(avl.GenerateDepthString(depth * length) + val)
	avl.GenerateBSTString(root.left, depth+1, "^", length)
}

func (avl *AVLTree) GenerateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString(" ")
	}
	return buffer.String()
}





