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

// Parent is used to get a pointer to the parent node of a given node
func (node *Node[T]) Parent() (*Node[T], error) {
	if node == nil {
		return nil, nodeNilError
	}
	return node.parent, nil
}

// LeftChild is used to get a pointer to the left child of a given node
func (node *Node[T]) LeftChild() (*Node[T], error) {
	if node == nil {
		return nil, nodeNilError
	}
	return node.left, nil
}

// RightChild is used to get a pointer to the right child of a given node
func (node *Node[T]) RightChild() (*Node[T], error) {
	if node == nil {
		return nil, nodeNilError
	}
	return node.right, nil
}

// BinarySearchTree struct will hold the core functionality of this library
type BinarySearchTree[T BinarySearchTreeElement] struct {
	root *Node[T]
}

// IsNil tells you if the pointer to the binary search tree is nil
func (bst *BinarySearchTree[T]) IsNil() bool {
	return bst == nil
}

// IsEmpty tells you if the binary search tree is empty
func (bst *BinarySearchTree[T]) IsEmpty() bool {
	return bst.IsNil() || bst.root == nil
}

// Root returns a pointer to the root of a binary search tree
func (bst *BinarySearchTree[T]) Root() *Node[T] {
	if bst.IsNil() {
		return nil
	}
	return bst.root
}
