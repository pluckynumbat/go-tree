// Package bstreelib: Basic Binary Search Tree Stuff
package bstreelib

import (
	"fmt"
)

var nodeNilError = fmt.Errorf("the node is nil")
var treeNilError = fmt.Errorf("the binary search tree is nil")
var treeEmptyError = fmt.Errorf("the binary search tree is empty")

// Node is the basic unit of the binary search tree, and contains data which can be anything that implements the comparable interface
type Node[T comparable] struct {
	data   T
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
}
