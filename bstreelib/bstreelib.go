// Package bstreelib: Basic Binary Search Tree Stuff
package bstreelib

import (
	"cmp"
	"fmt"
)

var nodeNilError = fmt.Errorf("the node is nil")
var treeNilError = fmt.Errorf("the binary search tree is nil")
var treeEmptyError = fmt.Errorf("the binary search tree is empty")

// BinarySearchTreeElement is a custom interface that combines the constraints of the Ordered and Stringer interfaces
type BinarySearchTreeElement interface {
	cmp.Ordered
	fmt.Stringer
}

// Node is the basic unit of the binary search tree, and contains data which can be anything that implements the BinarySearchTreeElement interface
type Node[T BinarySearchTreeElement] struct {
	data   T
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
}

// Node's implementation of the fmt.Stringer interface
func (node *Node[T]) String() string {
	if node == nil {
		return "nil"
	}
	return node.data.String()
}
