package BinarySearchTree

import (
	"fmt"
	"testing"
)

func comparable(i, j interface{}) int {
	return i.(int) - j.(int)
}

func TestBST_Add(t *testing.T) {


	bst := NewBST(comparable)
	bst.Add(7)
	bst.Add(4)
	bst.Add(9)
	bst.Add(2)
	bst.Add(5)
	bst.Add(8)
	bst.Add(11)
	bst.Add(1)
	bst.Add(3)
	bst.Add(12)
	bst.Add(10)
	fmt.Println(bst.String())
	fmt.Println()

	bst.PreOrderTraversal()
	fmt.Println()


}