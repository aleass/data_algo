package tree

import (
	"fmt"
	"testing"
)

func TestNode_PreOrder(t *testing.T) {
	tree := NewLinkedTree(33)
	tree.InsertOne(25)
	tree.InsertOne(51)
	tree.InsertOne(66)
	tree.InsertOne(19)
	tree.InsertOne(27)
	tree.InsertOne(16)
	tree.InsertOne(58)
	tree.InsertOne(34)
	tree.InsertOne(18)
	tree.InsertOne(13)
	tree.InsertOne(50)
	tree.InsertOne(17)
	tree.Del(25)
	fmt.Println(tree.Get(25))
}
