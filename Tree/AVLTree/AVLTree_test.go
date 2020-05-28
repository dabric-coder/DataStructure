package AVLTree

import (
	"fmt"
	"testing"
)

func comparable(i, j interface{}) int {
	return i.(int) - j.(int)
}


func TestAVLTree(t *testing.T) {
	avl := NewAVLNode(comparable)
	avl.Add(7)
	avl.Add(4)
	avl.Add(9)
	avl.Add(2)
	avl.Add(5)
	avl.Add(8)
	avl.Add(11)
	avl.Add(1)
	avl.Add(3)
	avl.Add(12)
	avl.Add(10)
	fmt.Println(avl.String())
	fmt.Println()

	avl.PreOrderTraversal()
	fmt.Println()

	avl.InOrderTraversal()
	fmt.Println()

	avl.PostOrderTraversal()
	fmt.Println()

	fmt.Println(avl.Predecessor(avl.root.left))
}
