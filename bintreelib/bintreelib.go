// Package bintreelib: Basic Binary Tree Stuff
package bintreelib

import (
	"fmt"

	"github.com/pluckynumbat/go-quez/sgquezlib"
)

var treeNilError = fmt.Errorf("the binary tree is nil")
var rootNilError = fmt.Errorf("the root is nil")

type Node struct {
	left  *Node
	data  string
	right *Node
}

// Node's implementation of the fmt.Stringer interface
func (node Node) String() string {
	return node.data
}

type BinaryTree struct {
	root *Node
}

